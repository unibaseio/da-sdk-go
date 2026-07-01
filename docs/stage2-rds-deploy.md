# Stage 2 deploy: shared Postgres index + read replicas

Goal: 1 writer + N read replicas sharing a Postgres (RDS) **index**, behind an
ALB. Only the gorm index (needle/bucket/conversation/account/volume/stat) moves
to Postgres — **blob content (logfs) stays local to the writer**. So this scales
list/get/`memoryStat` reads; content `download` on a replica falls back to the
DA network. Full content scaling is stage 3 (S3).

The code (branch `feat/rds-read-scaling`) is backward-compatible: with no
`HUB_DB_*` env it stays on local SQLite, unchanged.

---

## 0. Prerequisites

- RDS for PostgreSQL (15+), same VPC as the hubs, reachable on 5432 from the
  hub security group. Enable `rds.force_ssl` and use `sslmode=require`.
- A database + user, e.g. `hub` / `hubuser`.
- `pgloader` on the writer box (`sudo apt-get install -y pgloader`).
- A recent backup / EBS snapshot of `~/data/hub` before you start.

DSN used below (keep the password in an env/secret, not in files):
```
postgres://hubuser:PASSWORD@HUB-RDS.xxxx.ap-southeast-1.rds.amazonaws.com:5432/hub?sslmode=require
```

---

## 1. Migrate the SQLite index → Postgres

The needles table is ~34M rows, so load data WITHOUT indexes first, then let the
hub create the gorm indexes (incl. the covering `LOWER(owner),size` one).

### 1a. Freeze writes (consistency) and snapshot

```bash
cd ~/unibase-sdk-go
# stop the hub so gorm.db is quiescent (or take an EBS snapshot and migrate from a copy)
docker-compose -f docker-compose/docker-compose-hub.yml stop dimo-hub
cp -a ~/data/hub/gorm/gorm.db /tmp/gorm.migrate.db   # migrate from a static copy
```

### 1b. Bulk load with pgloader (tables + rows, no indexes)

Edit `scripts/sqlite-to-pg.load` (set the real RDS DSN), then:
```bash
time pgloader scripts/sqlite-to-pg.load
```
Rough expectation for 34M rows: tens of minutes (depends on RDS class + network).
`create no indexes` keeps the load fast; the hub builds indexes next.

### 1c. Add primary keys + reset sequences (pgloader's `create no indexes` dropped the PKs)

**Critical.** pgloader's `create no indexes` (and a CSV `\copy`) leave the tables
**without a PRIMARY KEY on `id`**. Without it every `ORDER BY id` — i.e. the
global (unscoped) list endpoints — degrades to a full seq scan + sort (the
unscoped `/api/listNeedle` measured ~5s on 34M rows). Add the PKs after the load,
**before** first start. Do the big table online (CONCURRENTLY), the small ones
inline:
```bash
PG="postgres://hubuser:PASSWORD@HUB-RDS...:5432/hub?sslmode=require"
# needles (34M): build the unique index online, then promote to PK
psql "$PG" -c "CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS needles_pkey ON needles (id);"
psql "$PG" -c "ALTER TABLE needles ADD CONSTRAINT needles_pkey PRIMARY KEY USING INDEX needles_pkey;"
# small tables: instant
psql "$PG" -c "ALTER TABLE accounts ADD PRIMARY KEY (id);
ALTER TABLE buckets ADD PRIMARY KEY (id);
ALTER TABLE volumes ADD PRIMARY KEY (id);
ALTER TABLE conversations ADD PRIMARY KEY (id);
ALTER TABLE stat_records ADD PRIMARY KEY (id);"
```
Then reset the id sequences to MAX(id) so new inserts don't collide with the
bulk-loaded rows:
```bash
psql "$PG" -c "DO \$\$ DECLARE t text; BEGIN
  FOR t IN SELECT unnest(ARRAY['accounts','buckets','needles','volumes','stat_records','conversations'])
  LOOP EXECUTE format('SELECT setval(pg_get_serial_sequence(%L,''id''), COALESCE((SELECT MAX(id) FROM %I),1))', t, t); END LOOP;
END \$\$;"
```
> The hub's `loadGORM` adds any missing PK idempotently as a backstop too — but
> doing it here (CONCURRENTLY for needles) avoids a multi-minute **blocking** PK
> build at first startup.

### 1d. Let the writer create the gorm schema + indexes

Point the writer at Postgres and start it once. On first connect it runs
`AutoMigrate` (reconciles schema) and `CREATE INDEX IF NOT EXISTS …` including
the partial covering `idx_needles_lower_owner_size_live` (`WHERE deleted_at IS
NULL`) — **this index build on 34M rows is the one-time, multi-minute,
write-locking step.**

Add to the writer's compose `environment:` and start:
```yaml
    environment:
      HUB_DB_DRIVER: "postgres"
      HUB_DB_DSN: "postgres://hubuser:PASSWORD@HUB-RDS...:5432/hub?sslmode=require"
```
```bash
docker-compose -f docker-compose/docker-compose-hub.yml rm -f dimo-hub
docker-compose -f docker-compose/docker-compose-hub.yml up -d dimo-hub
docker logs -f docker-compose_dimo-hub_1 2>&1 | grep -iE "gorm backend|memstat refreshed"
# expect: "gorm backend: postgres (shared index)" then later "memstat refreshed: ..."
```

After the index build, set the visibility map + planner stats so the partial
indexes serve **index-only** scans (a fresh bulk load leaves both stale, and the
`LOWER(owner)` expression indexes only get stats once they exist):
```bash
psql "$PG" -c "VACUUM (ANALYZE) needles;"
```

### 1e. Verify row counts match

```bash
SQLITE=/tmp/gorm.migrate.db PG="postgres://hubuser:PASSWORD@HUB-RDS...:5432/hub?sslmode=require" \
  scripts/verify-migration.sh
```
All tables must match. If they do, the writer is now serving from Postgres.

---

## 2. Deploy read replicas

Each replica: same image, same `HUB_DB_DSN`, plus `HUB_READONLY=1`. It serves
reads from the shared Postgres, skips all writes/DDL/chain-submit, and 503s any
upload that reaches it.

```yaml
    environment:
      HUB_DB_DRIVER: "postgres"
      HUB_DB_DSN: "postgres://hubuser:PASSWORD@HUB-RDS...:5432/hub?sslmode=require"
      HUB_READONLY: "1"
```
Launch as many replica instances as needed (separate EC2 / containers), all on
:8080.

> A replica's local `~/data/hub/logfs` is empty — so `download` of content it
> doesn't have falls back to the DA network. list/get/`memoryStat`/`memoryOverview`
> are served from the shared PG and scale across replicas.

---

## 3. ALB path-based routing (writes → writer, reads → any)

Target groups (all targets on :8080, health check `GET /api/info` expecting 200):

| Target group | Members |
|---|---|
| `tg-hub-writer` | the writer instance only |
| `tg-hub-read`   | the read replicas (optionally + writer) |

Listener rules on the HTTPS:443 listener, in priority order:

| Prio | Condition | Action |
|---|---|---|
| 10 | Path is `/api/upload`, `/api/uploadData`, `/api/memoryStat` **or** `/api/memoryOverview` | forward → `tg-hub-writer` |
| default | (everything else) | forward → `tg-hub-read` |

CLI sketch:
```bash
# writes + stats → writer
aws elbv2 create-rule --listener-arn $LISTENER --priority 10 \
  --conditions '[{"Field":"path-pattern","Values":["/api/upload","/api/uploadData","/api/memoryStat","/api/memoryOverview"]}]' \
  --actions   '[{"Type":"forward","TargetGroupArn":"'$TG_WRITER'"}]'
# default rule already forwards to tg-hub-read
```

Notes:
- `/api/memoryStat` + `/api/memoryOverview` go to the **writer**: the heavy
  recompute (full-index scan) runs writer-only, so replicas hold no snapshot and
  would return an empty (ComputedAt=0) result. Stats are low-QPS (dashboard), so
  routing them to the writer is fine.
- Match by path is enough (uploads are POST to those two paths). You can also add
  an `http-request-method = POST` condition for tightness (but keep the GET stat
  paths on the writer rule).
- Health check path `/api/info` is public (no auth) and cheap.
- If you put the writer in `tg-hub-read` too, reads also use it; to fully offload
  the writer, keep `tg-hub-read` = replicas only.

---

## 4. Rollback

- Revert the writer's env (remove `HUB_DB_*`) and restart → it goes back to local
  SQLite (`~/data/hub/gorm/gorm.db`), unchanged. The SQLite data was never
  modified by the migration (we copied it). Replicas are just turned off.
- The ALB can point everything back at the single writer target group.

---

## What this does / does not give you

- ✅ list / get reads scale across N replicas.
- ◑ `memoryStat` / `memoryOverview` are served by the **writer** (the recompute
  runs writer-only to avoid N× redundant full scans on the shared DB); they're
  low-QPS dashboard calls, so this is fine.
- ❌ `download` content does NOT scale (logfs is local to the writer; replicas
  fall back to the DA network). That needs stage 3 (S3).
- Writer remains the single writer (logfs/badger are still local single-writer).

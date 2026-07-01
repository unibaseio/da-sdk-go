package hub

import (
	"os"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// stage-2 RDS/Postgres integration tests. Skipped unless TEST_PG_DSN is set,
// e.g.:  TEST_PG_DSN="host=127.0.0.1 port=5433 user=hub password=hubpw dbname=hub sslmode=disable"
// Each test resets the hub tables first so they're order-independent.

func pgDSNOrSkip(t *testing.T) string {
	t.Helper()
	dsn := os.Getenv("TEST_PG_DSN")
	if dsn == "" {
		t.Skip("set TEST_PG_DSN to run the Postgres integration tests")
	}
	os.Setenv("HUB_DB_DRIVER", "postgres")
	os.Setenv("HUB_DB_DSN", dsn)
	os.Unsetenv("HUB_DB_DSN_READ")
	t.Cleanup(func() { os.Unsetenv("HUB_DB_DRIVER"); os.Unsetenv("HUB_DB_DSN"); os.Unsetenv("HUB_DB_DSN_READ") })
	// clean slate so tests don't pollute each other
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: glogger.Default.LogMode(glogger.Silent)})
	if err != nil {
		t.Fatalf("connect pg: %v", err)
	}
	db.Exec("DROP TABLE IF EXISTS accounts, buckets, needles, volumes, stat_records, conversations CASCADE")
	if sqldb, err := db.DB(); err == nil {
		sqldb.Close()
	}
	return dsn
}

// TestStage2PG: writer connect + AutoMigrate + partial LOWER(owner) indexes + PK,
// a LOWER(owner) write/read round-trip, dbresolver read/write-split registration,
// and readonly DDL-skip.
func TestStage2PG(t *testing.T) {
	dsn := pgDSNOrSkip(t)

	w := &Server{}
	w.loadGORM()
	if w.isSQLite() {
		t.Fatal("expected postgres backend, got sqlite")
	}
	for _, idx := range []string{
		"idx_needles_lower_owner_id_live", "idx_needles_lower_owner_size_live",
		"idx_needles_bucket", "idx_needles_name",
		"idx_volumes_lower_owner_file", "idx_buckets_lower_owner_id",
		"idx_conversations_lower_owner_id",
	} {
		var c int64
		w.gdb.Raw("SELECT count(*) FROM pg_indexes WHERE indexname = ?", idx).Scan(&c)
		if c != 1 {
			t.Errorf("index %s: want 1, got %d", idx, c)
		}
	}
	for _, dead := range []string{"idx_needles_owner", "idx_needles_owner_name"} {
		var c int64
		w.gdb.Raw("SELECT count(*) FROM pg_indexes WHERE indexname = ?", dead).Scan(&c)
		if c != 0 {
			t.Errorf("dead index %s should be dropped, got %d", dead, c)
		}
	}
	for _, tb := range []string{"accounts", "buckets", "needles", "volumes", "stat_records", "conversations"} {
		var c int64
		w.gdb.Raw("SELECT count(*) FROM pg_constraint WHERE conrelid = ?::regclass AND contype = 'p'", tb).Scan(&c)
		if c != 1 {
			t.Errorf("table %s: want 1 PK, got %d", tb, c)
		}
	}

	n := &types.Needle{Name: "n1", Bucket: "b1", Owner: "0xABC123", Size: 4096}
	if err := w.gdb.Create(n).Error; err != nil {
		t.Fatalf("insert needle: %v", err)
	}
	var got types.Needle
	if err := w.gdb.Where("LOWER(owner) = ?", "0xabc123").First(&got).Error; err != nil {
		t.Fatalf("read back by LOWER(owner): %v", err)
	}
	if got.Size != 4096 {
		t.Errorf("round-trip size: want 4096, got %d", got.Size)
	}

	// dbresolver read/write split (READ dsn = same PG here; validates registration + routing)
	os.Setenv("HUB_DB_DSN_READ", dsn)
	defer os.Unsetenv("HUB_DB_DSN_READ")
	r := &Server{}
	r.loadGORM()
	var rn types.Needle
	if err := r.gdb.Clauses(dbresolver.Read).Where("LOWER(owner) = ?", "0xabc123").First(&rn).Error; err != nil {
		t.Fatalf("dbresolver read-routed: %v", err)
	}
	if err := r.gdb.Clauses(dbresolver.Write).Where("LOWER(owner) = ?", "0xabc123").First(&rn).Error; err != nil {
		t.Fatalf("dbresolver write-routed read (gate-check pattern): %v", err)
	}

	// readonly must NOT run DDL
	os.Unsetenv("HUB_DB_DSN_READ")
	w.gdb.Exec("DROP INDEX IF EXISTS idx_conversations_name")
	ro := &Server{readonly: true}
	ro.loadGORM()
	var c int64
	ro.gdb.Raw("SELECT count(*) FROM pg_indexes WHERE indexname = 'idx_conversations_name'").Scan(&c)
	if c != 0 {
		t.Error("readonly loadGORM must not create indexes")
	}
	w.loadGORM()
	w.gdb.Raw("SELECT count(*) FROM pg_indexes WHERE indexname = 'idx_conversations_name'").Scan(&c)
	if c != 1 {
		t.Error("writer loadGORM should recreate idx_conversations_name")
	}
}

// TestStage2PG_MigrationBackstop simulates a pgloader bulk-load (rows present,
// but no PK and no indexes — pgloader's "create no indexes") and verifies the
// writer's loadGORM repairs it: adds PK + builds indexes, no row loss. This is
// the migration-path behavior docs/stage2-rds-deploy.md relies on.
func TestStage2PG_MigrationBackstop(t *testing.T) {
	pgDSNOrSkip(t)

	(&Server{}).loadGORM() // stand up the schema
	pg, _ := gorm.Open(postgres.Open(os.Getenv("HUB_DB_DSN")), &gorm.Config{Logger: glogger.Default.LogMode(glogger.Silent)})
	pg.Exec("DROP TABLE IF EXISTS needles CASCADE")
	pg.Exec(`CREATE TABLE needles (id bigint, created_at timestamptz, updated_at timestamptz, deleted_at timestamptz, owner text, bucket text, name text, size bigint)`)
	for i := 1; i <= 5; i++ {
		pg.Exec("INSERT INTO needles (id, owner, bucket, name, size) VALUES (?, ?, 'b', ?, ?)", i, "0xMig", "n", i*100)
	}

	w := &Server{}
	w.loadGORM() // must repair PK + indexes on the loaded table
	var pk, rows, idx int64
	w.gdb.Raw("SELECT count(*) FROM pg_constraint WHERE conrelid='needles'::regclass AND contype='p'").Scan(&pk)
	w.gdb.Raw("SELECT count(*) FROM needles").Scan(&rows)
	w.gdb.Raw("SELECT count(*) FROM pg_indexes WHERE indexname='idx_needles_lower_owner_id_live'").Scan(&idx)
	if pk != 1 || rows != 5 || idx != 1 {
		t.Errorf("migration backstop: pk=%d rows=%d idx=%d (want 1,5,1)", pk, rows, idx)
	}
}

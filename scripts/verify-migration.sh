#!/usr/bin/env bash
#
# Verify a SQLite → Postgres index migration by comparing per-table row counts.
#
# Usage:
#   SQLITE=/tmp/gorm.migrate.db \
#   PG="postgres://hubuser:PASSWORD@HUB-RDS...:5432/hub?sslmode=require" \
#   scripts/verify-migration.sh
#
# Requires: sqlite3, psql

set -u

: "${SQLITE:?set SQLITE=/path/to/gorm.db}"
: "${PG:?set PG=postgres://...}"

# gorm table names (struct → pluralized snake_case)
tables=(accounts buckets needles volumes conversations stat_records)

fail=0
printf '%-16s %14s %14s   %s\n' "table" "sqlite" "postgres" "status"
printf '%-16s %14s %14s   %s\n' "-----" "------" "--------" "------"
for t in "${tables[@]}"; do
  s=$(sqlite3 "$SQLITE" "SELECT count(*) FROM $t;" 2>/dev/null)
  p=$(psql "$PG" -tAc "SELECT count(*) FROM $t;" 2>/dev/null)
  s=${s:-ERR}; p=${p:-ERR}
  if [[ "$s" == "$p" && "$s" != "ERR" ]]; then
    status="✅"
  else
    status="❌ MISMATCH"
    fail=$((fail + 1))
  fi
  printf '%-16s %14s %14s   %s\n' "$t" "$s" "$p" "$status"
done

echo
if [[ "$fail" -eq 0 ]]; then
  echo "ALL TABLES MATCH ✅"
  exit 0
else
  echo "FAILED: $fail table(s) mismatched"
  exit 1
fi

#!/usr/bin/env bash
#
# Smoke-test the hub /api auth + owner-match + size-cap
# behaviour against a running hub-mock (or a real hub).
#
# Usage:
#   1) In one terminal, start the mock:
#        cd <repo root>
#        CHAIN_TYPE=bnb-testnet-dao go run ./cmd/hub-mock -addr 127.0.0.1:18086
#   2) In another terminal, run this script.
#
# Requirements:
#   - python3 + the membase package at $MEMBASE_REPO (default: ../membase) so we
#     can mint Authorization headers via build_authorization().
#   - curl, awk
#
# Override the target with HUB_BASE, e.g. HUB_BASE=https://testnet.hub.membase.io
# to test the real hub once it's deployed.

set -u

BASE="${HUB_BASE:-http://127.0.0.1:18086}"
MEMBASE_REPO="${MEMBASE_REPO:-$(cd "$(dirname "$0")/../../../../../membase" 2>/dev/null && pwd)}"

# canonical test key (NOT a real key). Matches:
#   addr = 0x6370eF2f4Db3611D657b90667De398a2Cc2a370C
SK="${MEMBASE_SECRET_KEY:-0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20}"
SIGNER="${MEMBASE_ACCOUNT:-0x6370eF2f4Db3611D657b90667De398a2Cc2a370C}"

if [[ -z "$MEMBASE_REPO" || ! -d "$MEMBASE_REPO/src/membase/storage" ]]; then
  echo "ERROR: membase repo not found. Set MEMBASE_REPO=/path/to/membase" >&2
  exit 2
fi

mint_header() {
  local label="$1"
  MEMBASE_SECRET_KEY="$SK" MEMBASE_ACCOUNT="$SIGNER" \
  python3 -c "
import sys
sys.path.insert(0, '$MEMBASE_REPO/src')
from membase.storage._auth import build_authorization
print(build_authorization(hash_label=b'$label'))
"
}

HDR_UP=$(mint_header upload)
HDR_DN=$(mint_header download)

div() { printf '\n── %s ──────────────────────────────────────\n' "$*"; }
expect() { # expect <wanted_code> <actual_code> <label>
  if [[ "$2" == "$1" ]]; then
    printf '  ✅ %s  →  HTTP %s\n' "$3" "$2"
  else
    printf '  ❌ %s  →  HTTP %s (wanted %s)\n' "$3" "$2" "$1"
    fails=$((fails + 1))
  fi
}

fails=0

div "T1  /api/info  bypass (no auth)"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" "$BASE/api/info")
expect 200 "$code" "/api/info open"

div "T2  /api/upload  no Authorization header"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/upload" \
  -H 'Content-Type: application/json' \
  -d '{"owner":"noah-2026","id":"hello","message":"hi"}')
expect 401 "$code" "missing header rejected"

div "T3  /api/upload  owner = 'noah-2026' (not a 0x addr)"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/upload" \
  -H 'Content-Type: application/json' \
  -H "Authorization: $HDR_UP" \
  -d '{"owner":"noah-2026","id":"hello","message":"hi"}')
expect 401 "$code" "non-0x owner rejected"

div "T4  /api/upload  owner = some other ETH addr"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/upload" \
  -H 'Content-Type: application/json' \
  -H "Authorization: $HDR_UP" \
  -d '{"owner":"0x0000000000000000000000000000000000000001","id":"x","message":"y"}')
expect 401 "$code" "mismatched owner rejected"

div "T5  /api/upload  owner == signer  (happy path)"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/upload" \
  -H 'Content-Type: application/json' \
  -H "Authorization: $HDR_UP" \
  -d "{\"owner\":\"$SIGNER\",\"id\":\"hello\",\"message\":\"hi\"}")
expect 200 "$code" "valid upload accepted"

div "T6  /api/download  no owner → defaults to signer"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/download" \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H "Authorization: $HDR_DN" \
  --data-urlencode 'id=hello')
expect 200 "$code" "download owner-default"

div "T7  /api/download  owner = other  → reject"
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/download" \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H "Authorization: $HDR_DN" \
  --data-urlencode 'owner=0x000000000000000000000000000000000000dead' \
  --data-urlencode 'id=hello')
expect 401 "$code" "mismatched download owner rejected"

div "T8  body size cap (5 MiB, default cap 4 MiB)"
python3 -c "
import sys
body = '{\"owner\":\"$SIGNER\",\"id\":\"x\",\"message\":\"' + ('a' * (5*1024*1024)) + '\"}'
sys.stdout.buffer.write(body.encode())
" > /tmp/big.json
code=$(curl -s -o /tmp/r.json -w "%{http_code}" -X POST "$BASE/api/upload" \
  -H 'Content-Type: application/json' \
  -H "Authorization: $HDR_UP" \
  --data-binary @/tmp/big.json)
if [[ "$code" != "200" ]]; then
  printf '  ✅ oversized body rejected  →  HTTP %s\n' "$code"
else
  printf '  ❌ oversized body accepted (200)\n'
  fails=$((fails + 1))
fi
rm -f /tmp/big.json /tmp/r.json

echo
if [[ "$fails" -eq 0 ]]; then
  echo "ALL GREEN ✅"
  exit 0
else
  echo "FAILED: $fails check(s)"
  exit 1
fi

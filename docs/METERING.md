# DA Hub Metering

Metering is an accounting layer around successful hub API usage. It estimates
write cost, rejects writes a user cannot pay for, records usage in a persistent
ledger, accumulates unsettled fees per wallet, and settles those fees either
off-chain or through ERC-8183. It does not replace or alter DA storage, needle
metadata, or piece/proof submission.

Metering is disabled by default. With `HUB_METERING_ENABLED=false` all existing
upload and download behavior is unchanged.

## Configuration

All settings are environment variables. Amounts are token base units (wei) as
base-10 integer strings; floats are never used.

```
HUB_METERING_ENABLED=false            # master switch
HUB_METERING_CHARGE_WRITES=true       # record/charge writes when enabled
HUB_METERING_CHARGE_READS=false       # public downloads are not charged yet

HUB_METERING_WRITE_BASE_WEI=0         # flat fee per write
HUB_METERING_WRITE_PER_KB_WEI=0       # fee per billable KB (ceil)
HUB_METERING_READ_PER_REQUEST_WEI=0   # flat fee per read

HUB_METERING_DEFAULT_CREDIT_LIMIT_WEI=0   # 0 = unlimited local credit
HUB_METERING_CHECK_CHAIN=false            # consult chain balance/allowance
HUB_METERING_SETTLEMENT_MODE=offchain     # offchain | erc8183

HUB_METERING_AUTO_SETTLE=false        # run the background settlement worker
HUB_METERING_SETTLE_INTERVAL_SEC=300  # worker scan interval
HUB_METERING_SETTLE_THRESHOLD_WEI=0   # only settle accounts at/above this

HUB_PROVIDER_ADDRESS=                 # provider/spender address
HUB_PROVIDER_PRIVATE_KEY=             # signs transferFrom + ERC-8183 txs
HUB_ERC20_TOKEN_ADDR=                 # payment token
HUB_ERC8183_CONTRACT_ADDR=            # escrow contract
HUB_ERC8183_EVALUATOR_ADDR=           # evaluator / hook
HUB_CHAIN_RPC_URL=
HUB_CHAIN_ID=97
```

Notes:

- `HUB_METERING_DEFAULT_CREDIT_LIMIT_WEI=0` means unlimited local credit unless
  chain checks are enabled.
- `HUB_METERING_CHECK_CHAIN=false` means no RPC call is made during upload
  preflight.
- `HUB_METERING_SETTLEMENT_MODE=offchain` clears local debt only. Use `erc8183`
  once the ledger is stable and the provider wallet is funded for gas.

## Pricing model

```
write_fee = write_base_wei + ceil(payload_bytes / 1024) * write_per_kb_wei
read_fee  = read_per_request_wei
```

Bytes are rounded up to the next whole KB. Example: 2500 bytes bills 3 KB, so
`write_fee = write_base_wei + 3 * write_per_kb_wei`. All arithmetic uses
`math/big.Int`; the ledger stores decimal strings.

## Approve / deposit expectations

Users do not sign a transaction per upload. The intended flow:

1. Hold the payment token.
2. Approve the provider address to spend up to a chosen allowance.
3. Upload normally until allowance, balance, or credit limit is exhausted.
4. Settle (or let the worker settle) to clear debt and continue.

When `HUB_METERING_CHECK_CHAIN=true`, the hub reads on-chain balance and
allowance during write preflight and rejects writes that exceed either.

## Upload behavior

On every write the hub estimates the fee and runs an admission check before
touching storage:

```
required = unsettled_fee + estimated_fee
reject if account disabled
reject if credit_limit > 0 and required > credit_limit
if chain checks on: reject if balance < required or allowance < required
```

A write is charged only after storage succeeds. Failed writes create no billing
event. When metering is disabled or write charging is off, uploads behave
exactly as before.

Public downloads are not charged in this version. They have no requester
identity; requester-paid reads would require authenticated downloads.

## HTTP endpoints

Public (no auth):

```
GET /api/metering/pricing
GET /api/metering/usage?owner=0x...
GET /api/metering/can-write?owner=0x...&bytes=123
```

`can-write` always returns 200 with a decision body; it is a preflight query,
not a write attempt.

Authenticated (signed `Authorization`, signer must equal owner; owner defaults
to the signer if omitted):

```
POST /api/metering/settle
```

## 402 handling

When a real upload is refused, the hub returns `402 Payment Required` with:

```json
{
  "allowed": false,
  "reason": "insufficient_allowance",
  "required_wei": "1300000000000000",
  "estimated_fee_wei": "1300000000000000",
  "unsettled_fee_wei": "5000000000000000",
  "credit_limit_wei": "1000000000000000000",
  "balance_wei": "4000000000000000",
  "allowance_wei": "0",
  "action": "approve_or_settle"
}
```

Reasons: `metering_disabled`, `account_disabled`, `credit_limit_exceeded`,
`insufficient_allowance`, `insufficient_balance`, `chain_check_failed`,
`allowed`. Chain/config problems surface as `chain_check_failed` with a 402, not
a 500.

## Settlement behavior

Off-chain (`SETTLEMENT_MODE=offchain`): a single DB transaction sums the
unsettled events, records a confirmed settlement, marks the events settled, and
subtracts the amount from the account.

ERC-8183 (`SETTLEMENT_MODE=erc8183`): the on-chain sequence is

```
ERC20.transferFrom(user -> provider, amount)
ERC8183.createJob(...)
ERC8183.setBudget(jobId, token, amount)
ERC8183.fund(jobId, amount)
ERC8183.submit(jobId, reportHash)
```

The deliverable is `sha256` of a canonical settlement report:

```json
{
  "amount_wei": "...",
  "bytes_written": 123456,
  "from_event_id": 10,
  "owner": "0x...",
  "reads": 0,
  "timestamp": 1234567890,
  "to_event_id": 50,
  "type": "da-hub-metering-settlement",
  "writes": 12
}
```

Debt is cleared only after all chain calls succeed. The settlement response and
record hold the transfer / create-job / fund / submit tx hashes, the job id, and
the report hash.

## Retry behavior

ERC-8183 settlement reserves the covered events (status `settling`) before the
chain calls. On failure the events are reverted to `unsettled` and the account's
unsettled fee is left intact, so the debt is never lost and the next settlement
attempt (manual or worker) retries it. The failed settlement record keeps the
error and any tx hashes produced before the failure.

The critical invariant: never clear unsettled debt before chain settlement
succeeds.

## Background worker

With `HUB_METERING_AUTO_SETTLE=true` a worker scans every
`HUB_METERING_SETTLE_INTERVAL_SEC` seconds, selecting enabled accounts whose
unsettled fee is at or above `HUB_METERING_SETTLE_THRESHOLD_WEI` and settling
each. A failure or panic settling one account is logged and never aborts the
scan or crashes the hub. The worker stops during hub shutdown.

This version assumes a single hub instance. For multi-replica deployment, add DB
locking or leader election before enabling the worker in production.

## Operational metrics

Start with logs; Prometheus counters can follow. Signals worth tracking:

```
record_write_total / record_read_total
rejected_write_total
unsettled_fee_total
settlement_success_total / settlement_failed_total
chain_check_failed_total
insufficient_allowance_total / insufficient_balance_total
```

Alerts worth wiring:

```
provider wallet gas low
provider private key missing while erc8183 enabled
RPC errors rising
settlement failures rising
unsettled fee total too high
unexpected 402 spike
```


# Unibase DA

**Unibase DA** is a high-performance, decentralized data availability (DA) and storage layer optimized for AI and DePIN workloads. Built with zk-based on-chain verification, it supports massive data throughput and EB-scale capacity while ensuring trust and transparency via Ethereum-based settlement.

---

## 🔍 Why Unibase DA?

Current DA solutions (EigenDA, Celestia, Avail, etc.) were designed for rollup transaction data, not the large-scale, real-time data workloads required by AI and DePIN. Unibase DA solves this by offering:

- 🔐 On-chain verifiability with ZK Proofs
- ⚡️ 100GB/s write throughput & 100MB/s encoding
- 🌐 Horizontal scaling to millions of devices
- 🧠 Native support for AI/DePIN memory + model data
- 📦 Modular storage APIs for files, models, blobs, and streams

---

## 📊 DA Protocol Comparison

| Feature                      | Ethereum | EigenDA | Celestia | Avail | 0G | **Unibase DA** |
|-----------------------------|----------|---------|----------|-------|----|----------------|
| **Data Verification**       | —        | Off-chain | Off-chain | Off-chain | Off-chain | ✅ **On-chain (ZK Proofs)** |
| **Data Security Model**     | Honest Majority | Honest Majority | Honest Majority | Honest Majority | Honest Majority | ✅ **Honest One (On-chain)** |
| **Data Throughput**         | 0.083 MB/s | 10 MB/s | 1.5–10 MB/s | — | 50 GB/s | ✅ **100 GB/s** |
| **Consensus**               | Casper   | None    | Tendermint | BABE + GRANDPA | Tendermint | None |
| **DAS Support**             | ❌        | ❌      | ✅        | ✅     | ✅  | ✅ |
| **Proof Type**              | Validity Proof | Validity Proof | Fraud Proof | Validity Proof | Validity Proof | ✅ **Fraud Proof** |
| **DA Settlement Layer**     | Ethereum | Ethereum | Celestia | Avail | 0G | ✅ **Base / Ethereum** |

---

## 📦 Features

- **Write Throughput:** 100 GB/s+
- **Encoding Speed:** 100 MB/s
- **Low-Latency Access:** Built for real-time AI/DePIN scenarios

### ✅ High Scalability

- **Capacity:** EB+
- **Nodes:** 1M+ edge/DePIN nodes
- **Storage Pools:** Public & private options

### ✅ Programmable

- **Ownership & Access Control:** On-chain via smart contracts
- **ZK Proofs:** Enable verifiable data integrity
- **Assetization:** Built-in support for data NFTs & streams

---

## 🧠 Architecture

Unibase DA follows a multi-layer architecture to ensure speed, reliability, and transparency:

- **Smart Contract Layer:** Receives commitment, validates encode proofs, settles data
- **Storage Nodes:** Store encoded data blocks
- **ZK Verifier:** Validates submitted proofs
- **Client SDK:** Uploads files, generates proofs, interacts with contracts

![Unibase DA Verification Diagram](https://path-to-image/da-verification.png)

---

## 📦 Installation

```bash
git clone https://github.com/unibaseio/da-sdk-go.git
cd da-sdk-go
go build ./...
```

---

## 📚 Quick Usage

Two access modes — pick by workload:

- **Direct** (`ubcli` / `example/*`, recommended for one-off files & dirs): the client
  discovers a stream via the gateway (or pins one), uploads **straight to it**, and the
  client's **own wallet** submits the on-chain commitment. Fewest hops; the uploader owns
  the piece.
- **Hub** (a lightweight service exposing the S3-shaped `/v1` object-store API): best for
  **continuous / many small writes** (e.g. AI agent memory). It batches writes and serves a
  `bucket → object` API. See “Hub node” below.

### Supported chains

Set `CHAIN_TYPE` to one of: `base-sepolia` (default), `base`, `bsc-mainnet`,
`eth-mainnet`, `bnb-testnet-v2`, `bnb-testnet-dao`, `local-anvil`.
Override the RPC with `CHAIN_RPC_<chainID>` / `CHAIN_RPC_FILTER_<chainID>`
(generic `CHAIN_RPC` works for single-chain setups). Gas knobs:
`GAS_LIMIT`, `GAS_TIP` (EIP-1559 priority fee, wei), `GAS_PRICE` (legacy /
zero-baseFee chains).

### `ubcli` — the CLI (direct upload/download)

```bash
make cli                       # builds bin/ubcli (native)
export UNIBASE_GATEWAY=https://<gateway>     # stream discovery + metadata
export UNIBASE_KEY=<hex key>                 # upload: funded wallet (ETH gas + UB bond)

# upload a file or a whole directory (each file → its own on-chain commitment)
ubcli da upload   --path ./weights.bin [--stream 0x<streamAddr>] [--name weights.bin]
# download by name (reconstructed from DA pieces; reads sign with any key, no funds)
ubcli da download --name weights.bin --out ./weights.bin
ubcli da ls                                  # list files on the gateway
```

`--stream` pins a specific stream node; omit it and the gateway picks an online one.
Add `--json` to any command for machine-readable output (agent-friendly). Reads are public.

### SDK examples (direct)

```bash
export CHAIN_TYPE=base-sepolia
# upload a file/dir straight to a stream + client-side AddPiece
go run ./example/upload   -gateway https://<gateway> -sk <key> -path <file|dir> [-stream 0x<addr>] [-name <name>]
# reconstruct a file from its DA pieces
go run ./example/download -gateway https://<gateway> -name <file_name> -out <save_path>
```

### ZK params

Proof generation and verification need the shared ceremony SRS (~7.6 GB, cached
in `~/.plonk`, sha256-verified). Fetch programmatically via `param.FetchAll()`
(node daemons do this automatically on first start). A locally generated SRS
will not verify against the on-chain `KZG_VK`.

---

## 🛰 Hub node — verifiable object store (`/v1`)

The hub is a lightweight service that turns Unibase DA into an **S3-shaped, wallet-native
object store**: `owner → bucket{kind} → object{key} → DA piece(s)`. It's the right choice
for **continuous / small writes** (agent memory, RAG chunks) — it batches writes and serves
listing/download — whereas one-off large files are better via the direct CLI above.

```bash
make hub               # builds bin/hub-edge (GOOS=linux; use `go build ./app/hub` for the host OS)
./hub-edge init
./hub-edge daemon run -e http://<public-ip>:8084
```

### `/v1` API (resource-oriented)

```
PUT    /v1/buckets/{bucket}                 # declare a bucket + kind (idempotent; kind immutable)
GET    /v1/buckets [?owner=&kind=&cursor=]  # list buckets
PUT    /v1/buckets/{bucket}/objects/{key}   # put one object (raw body); ?wait=1 blocks until on-chain
POST   /v1/buckets/{bucket}/objects         # batch put (multipart files[])
GET    /v1/buckets/{bucket}/objects/{key}[/content|/proof]   # metadata / bytes / verification bundle
GET    /v1/stats | /v1/overview | /v1/owners/{owner}
```

`kind ∈ {memory, knowledgebase, model, dataset, file}` — a bucket property (declared once)
that selects the storage profile; you don't pass it per upload. Every write returns a
**verifiable receipt**: `{commitment, chain:{txHash}, status: staged→committed, availability}`.

**Auth:** writes require a signed `Authorization` header (`sdk.DecodeAuth` format, ±10 min
window) and enforce **owner == signer**. Reads are **public** — stored content is
client-encrypted, so listing/download exposes only ciphertext + metadata. The legacy `/api/*`
surface stays for existing clients. Tunables:

| Variable | Default | Purpose |
|----------|---------|---------|
| `HUB_DB_DRIVER` / `HUB_DB_DSN` | sqlite | shared Postgres index for multi-instance read scaling |
| `HUB_READONLY` | unset | run as a read-only replica (no writes/DDL/chain submit) |
| `HUB_AUTH_DRIFT_SEC` | `600` | auth timestamp window |
| `HUB_MAX_JSON_BYTES` / `HUB_MAX_MULTIPART_BYTES` | 4 MiB / 64 MiB | body caps |
| `HUB_RATE_IP_RPS` / `HUB_RATE_OWNER_RPS` | `1000` / `1000` | per-IP / per-owner rate limits |

For local API smoke-testing without a chain or disk, use the standalone mock:
`go run ./cmd/hub-mock`.

---

## 🧱 Contract deployment (`contract/deploy`)

Deploys the full V2 proxy suite (verifiers → implementations → ERC-1967 proxies →
cross-wiring → VK roots → min pledges → base penalty):

```bash
go run ./contract/deploy -sk <deployer-key> \
  -rpc <rpc-url> -chainid <id> \
  [-skip-dao] [-slots <blocks-per-epoch>] [-min-prove-time <blocks>]
```

Notes from the local-Anvil integration run:

- Addresses are deterministic in deploy order; the `LocalAnvil` table in
  `contract/common/common.go` matches the current sequence — don't reorder.
- The tool sets `basePenalty` to the production value (**10000 UB**,
  `contract/common.DefaultPenalty`) post-deploy via the governance setter, and
  `minPledge(type 1) = 5× penalty` (concurrent-challenge capacity = 5) so stores
  always have enough locked stake to be challengeable. Order matters: the min
  pledge is set first — `setBasePenalty` in current `da-contract` requires
  `minStakeOf(1) >= penalty`.
- ⚠️ The bindings under `contract/v2/go` embed the **bytecode that gets
  deployed**. After changing `da-contract` impls, regenerate the affected
  bindings (forge artifact → abigen) before deploying/upgrading, or the tool
  ships stale contracts.
- After any redeploy on a public chain, update the per-chain address table in
  `contract/common/common.go` (and the chain's `KZGVKRoot` if the SRS changed).

The end-to-end fraud-proof test suite lives in `da-go/contract/test-v2`
(see its README for the Anvil runbook).

---

## 🔗 Links

- 🌐 Website: [https://www.unibase.com](https://www.unibase.com)
- 📖 GitBook: [Unibase Docs](https://openos-labs.gitbook.io/unibase-docs/)
- 🧠 GitHub: [github.com/unibaseio](https://github.com/unibaseio)
- 🐦 Twitter: [@Unibase_AI](https://twitter.com/Unibase_AI)
- 📢 Telegram: [@unibase_ai](https://t.me/unibase_ai)

---

## 📜 License

MIT License

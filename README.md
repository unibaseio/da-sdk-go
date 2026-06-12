
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
| **DA Settlement Layer**     | Ethereum | Ethereum | Celestia | Avail | 0G | ✅ **Base / BSC / Ethereum** |

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

## 🚀 SDK Quickstart

```bash

## 📦 Installation

```bash
git clone https://github.com/unibaseio/da-sdk-go.git
cd da-sdk-go
go build ./...
```

---

## 📚 Quick Usage

### Supported chains

Set `CHAIN_TYPE` to one of: `base-sepolia` (default), `base`, `bsc-mainnet`,
`eth-mainnet`, `bnb-testnet-v2`, `bnb-testnet-dao`, `local-anvil`.
Override the RPC with `CHAIN_RPC_<chainID>` / `CHAIN_RPC_FILTER_<chainID>`
(generic `CHAIN_RPC` works for single-chain setups). Gas knobs:
`GAS_LIMIT`, `GAS_TIP` (EIP-1559 priority fee, wei), `GAS_PRICE` (legacy /
zero-baseFee chains).

### Upload a file/directory

```bash
export CHAIN_TYPE=bnb-testnet-dao
cd example/upload
go build
./upload --model=false --sk=<your_secret_key> --path=<your_local_path>
```

### Download a file/directory

```bash
export CHAIN_TYPE=bnb-testnet-dao
cd example/download
go build
./download --model=false --sk=<your_secret_key> --name=<file_name> --path=<your_save_path>
```

### ZK params

Proof generation and verification need the shared ceremony SRS (~7.6 GB, cached
in `~/.plonk`, sha256-verified). Fetch programmatically via `param.FetchAll()`
(node daemons do this automatically on first start). A locally generated SRS
will not verify against the on-chain `KZG_VK`.

---

## 🛰 Hub node

The hub is a lightweight HTTP gateway for upload/download without running a full
node:

```bash
make hub               # builds bin/hub-edge (GOOS=linux; use `go build ./app/hub` for the host OS)
./hub-edge init
./hub-edge daemon run -e http://<public-ip>:8084
```

All `/api` routes (except `/api/info`) require a signed `Authorization` header
(`sdk.DecodeAuth` format, ±10 min timestamp window) and enforce
**owner == signer** on reads and writes. Tunables:

| Variable | Default | Purpose |
|----------|---------|---------|
| `HUB_AUTH_DRIFT_SEC` | `600` | auth timestamp window |
| `HUB_MAX_JSON_BYTES` | `4194304` (4 MiB) | JSON body cap |
| `HUB_MAX_MULTIPART_BYTES` | `67108864` (64 MiB) | multipart body cap |
| `HUB_RATE_IP_RPS` / `HUB_RATE_IP_BURST` | `10` / burst | per-IP rate limit |
| `HUB_RATE_OWNER_RPS` / `HUB_RATE_OWNER_BURST` | `5` / burst | per-owner rate limit |

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
- The proxies initialize `basePenalty` to 1e18; the tool sets the production
  value (**10000 UB**, `contract/common.DefaultPenalty`) post-deploy via the
  governance setter, and `minPledge(type 1) = 2× penalty` so stores always have
  enough locked stake to be challengeable.
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

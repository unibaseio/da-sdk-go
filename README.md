# Unibase DA SDK (Go)

The official Go SDK and client library for **Unibase DA** — a decentralized data
availability & storage layer with **on-chain ZK verification** under a fraud-proof
(“honest-one”) model, settling on Ethereum L2s (Base) and Ethereum.

This repository is the developer entry point: use it to **upload and download data**,
talk to a **hub** (a lightweight S3-shaped object gateway), and interact with the
on-chain DA contracts — without running your own node.

- Node runtime (gateway / store / stream / validator / hub): [`da-go`](https://github.com/unibaseio/da-go)
- On-chain contracts (Foundry): [`da-contract`](https://github.com/unibaseio/da-contract)
- ZK proof system: `da-core`

---

## Install

```bash
go get github.com/unibaseio/da-sdk-go
```

Requires Go 1.22+. The crypto library links [blst](https://github.com/supranational/blst),
so builds are **CGO-enabled** (a C toolchain must be available).

---

## Concepts (30 seconds)

Data is **content-addressed** and organized in three levels:

```
File  ──►  Piece  ──►  Replica
 (your      (a slice,   (one erasure-coded shard,
  upload)    ≤ ~1GB)     stored on one node)
```

- A file is chunked into **pieces**; each piece is **Reed–Solomon (N,K)** erasure-coded
  into `N` replicas so any `K` can rebuild it.
- Every object’s name is the hex of its cryptographic commitment — **immutable and
  self-verifying**. Correct encoding is proven **on-chain via ZK**, and storage is kept
  honest by a **challenge/response fraud game** (a single honest party can force
  correctness).
- Supported erasure policies (N/K): `6/4`, `14/7`, `32/16`, `64/32`.

You don’t need to know the proof internals to use the SDK — uploads return a receipt,
downloads verify automatically.

---

## Quick start — CLI (`ubcli`)

The fastest way to push/pull data. Build it:

```bash
make cli          # produces bin/ubcli
```

```bash
# configure the signing key + gateway endpoint (see `ubcli --help`)
ubcli key ...                     # manage the signing key
ubcli da upload   --path ./model.bin --name my-model --n 6 --k 4
ubcli da download --name my-model --path ./out.bin
```

`ubcli` is agent/script friendly: it signs requests with your key, talks to a gateway,
and prints JSON. Run `ubcli <command> --help` for the exact flags.

---

## Quick start — Go SDK

The `sdk` package is a thin, typed client over the gateway/hub HTTP API. Key calls:

| Task | Function |
|------|----------|
| Upload bytes / a file | `sdk.Upload`, `sdk.UploadData`, `sdk.UploadFileMeta` |
| Download & reconstruct | `sdk.Download`, `sdk.DownloadParallel`, `sdk.DownloadPiece`, `sdk.DownloadWSize` (range) |
| Fetch receipts / metadata | `sdk.GetFileReceipt`, `sdk.GetPieceReceipt`, `sdk.GetReplicaReceipt` |
| Node registry | `sdk.GetEdge`, `sdk.EnsureEdge` |
| Auth | `sdk.Login`, `sdk.DecodeAuth` / `sdk.VerifyAuth` |

Downloads are **trustless**: the client re-derives the commitment from the returned
bytes and checks it against the on-chain receipt before handing data back. Missing data
shards are rebuilt from surviving replicas via the erasure code.

The on-chain client lives in `contract/v2` (`ContractManage`) — `AddPiece`, `AddReplica`,
epoch proofs, and the challenge/response game — with generated bindings under
`contract/v2/go`.

---

## Hub client (S3-shaped `/v1`)

A **hub** is a lightweight, wallet-native object gateway: a clean, S3-shaped `/v1` API
(buckets/objects) backed by verifiable DA storage. Use it as a drop-in decentralized
object store for AI memory, models, and datasets. See `hub/` and `sdk/hub.go`.

```bash
make hub          # build the hub server binary
```

Auth is signature-based (an `Authorization` header signed by your wallet, within a short
time window); every namespace is owner-scoped.

---

## Build

```bash
make cli          # ubcli
make hub          # hub server
make build        # library + tooling
make test         # unit tests   (make test-slow for the heavy crypto suites)
```

ZK proving/verification parameters (the KZG structured reference string, etc.) are
**fetched on demand** from a shared, SHA256-verified distribution into `~/.plonk` — you
do not generate them locally.

---

## Repository layout

| Path | What |
|------|------|
| `sdk/` | Developer API (upload/download, receipts, hub, auth) |
| `lib/bls/` | Crypto primitives (KZG over BLS12-377, erasure coding, commitments) |
| `lib/types/` | Wire/storage data model (`File` / `Piece` / `Replica`), CBOR |
| `lib/` | Infra: `kv` (BadgerDB), `logfs`, `merkle`, `key` (keystore), `piece`, `config` |
| `contract/v2/` | On-chain client + generated bindings (`go/`) + deploy tool (`deploy/`) |
| `hub/` | Lightweight S3-shaped object gateway |
| `param/` | SRS / proving-key fetch |
| `estimator/` | L1 fee estimation |
| `app/` | `cli` (ubcli), `hub`, `cmd` |

---

## Status & contributing

Unibase DA is under active development; APIs may change ahead of a `v1` tag. Issues and
PRs are welcome — please open an issue to discuss substantial changes first.

Docs: <https://unibaseio.gitbook.io/unibase-docs/> · Explorer: <https://www.explorer.unibase.com/>

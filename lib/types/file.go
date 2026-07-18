package types

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

// CurrentWireVersion tags the CBOR wire/persistence format of the core records
// below (V6 / FORMAT_VERSIONING_DESIGN.md). fxamacker encodes by field name, so
// this field is additive-safe: legacy blobs have no Version key → decode as 0,
// treated as v1. Bump ONLY on a breaking change to a record's meaning, and give
// readers a matching decode path. Version==0 means "pre-versioning (v1)".
const CurrentWireVersion uint16 = 1

// stampVersion sets the current wire version on a fresh (0) record before
// marshaling; an already-set version is preserved.
func stampVersion(v *uint16) {
	if *v == 0 {
		*v = CurrentWireVersion
	}
}

// checkVersion fails closed on a record written by a newer node than we
// understand — never silently misparse a future layout as the current one.
func checkVersion(v uint16) error {
	if v > CurrentWireVersion {
		return fmt.Errorf("unsupported wire version %d (this node understands <= %d) — upgrade the node", v, CurrentWireVersion)
	}
	return nil
}

type Policy struct {
	N uint8
	K uint8
}

// SupportedPolicies is the single source of truth for the RS (N,K) codes the
// protocol supports. Check validates against it, and the deploy tool iterates
// it to register per-policy VK roots — keep new policies in one place.
var SupportedPolicies = []Policy{
	{N: 6, K: 4},
	{N: 14, K: 7},
	{N: 32, K: 16},
	{N: 64, K: 32},
}

func (p Policy) Check() error {
	for _, sp := range SupportedPolicies {
		if p.N == sp.N && p.K == sp.K {
			return nil
		}
	}
	return fmt.Errorf("unsupported rs policy: %d %d", p.N, p.K)
}

type FileCore struct {
	Version  uint16
	Policy   Policy
	Name     string
	Hash     string
	Size     int64
	Owner    common.Address
	Creation time.Time
}

func (frc *FileCore) Serialize() ([]byte, error) {
	stampVersion(&frc.Version)
	return cbor.Marshal(frc)
}

func (frc *FileCore) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, frc); err != nil {
		return err
	}
	return checkVersion(frc.Version)
}

type FileReceipt struct {
	FileCore
	ChainType string
	Pieces    []string
}

func (fr *FileReceipt) Serialize() ([]byte, error) {
	stampVersion(&fr.Version)
	return cbor.Marshal(fr)
}

func (fr *FileReceipt) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, fr); err != nil {
		return err
	}
	return checkVersion(fr.Version)
}

type FileFull struct {
	FileReceipt
	Proofs     [][]byte
	PieceSizes []int64
}

func (ff *FileFull) Serialize() ([]byte, error) {
	stampVersion(&ff.Version)
	return cbor.Marshal(ff)
}

func (ff *FileFull) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, ff); err != nil {
		return err
	}
	return checkVersion(ff.Version)
}

type PieceCore struct {
	Version  uint16
	Policy   Policy
	Name     string
	Serial   uint64
	Size     int64 // raw size
	Start    uint64
	Expire   uint64
	Price    *big.Int
	Owner    common.Address
	Streamer common.Address
	TxHash   string
}

func (pc *PieceCore) Serialize() ([]byte, error) {
	stampVersion(&pc.Version)
	return cbor.Marshal(pc)
}

func (pc *PieceCore) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, pc); err != nil {
		return err
	}
	return checkVersion(pc.Version)
}

type PieceReceipt struct {
	PieceCore
	Creation  time.Time
	ChainType string
	Replicas  []string
	StoredOn  []common.Address
}

func (cr *PieceReceipt) Serialize() ([]byte, error) {
	stampVersion(&cr.Version)
	return cbor.Marshal(cr)
}

func (cr *PieceReceipt) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, cr); err != nil {
		return err
	}
	return checkVersion(cr.Version)
}

type PieceWitness struct {
	Choose  uint8
	Witness []byte
}

type ReplicaCore struct {
	Version  uint16
	Fake     bool
	Name     string // encoded
	Serial   uint64
	Size     int64  // stored size
	Piece    string // belongs to which piece
	Index    uint8  // index in piece
	StoredOn common.Address
	Ordinal  uint64 // index of store
	TxHash   string
}

func (rc *ReplicaCore) Serialize() ([]byte, error) {
	stampVersion(&rc.Version)
	return cbor.Marshal(rc)
}

func (rc *ReplicaCore) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, rc); err != nil {
		return err
	}
	return checkVersion(rc.Version)
}

type ReplicaReceipt struct {
	ReplicaCore
	Creation  time.Time
	ChainType string
}

func (rr *ReplicaReceipt) Serialize() ([]byte, error) {
	stampVersion(&rr.Version)
	return cbor.Marshal(rr)
}

func (rr *ReplicaReceipt) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, rr); err != nil {
		return err
	}
	return checkVersion(rr.Version)
}

type ReplicaWitness struct {
	Version uint16
	Index   uint64
	Proof   []byte
}

func (rw *ReplicaWitness) Serialize() ([]byte, error) {
	stampVersion(&rw.Version)
	return cbor.Marshal(rw)
}

func (rw *ReplicaWitness) Deserialize(b []byte) error {
	if err := cbor.Unmarshal(b, rw); err != nil {
		return err
	}
	return checkVersion(rw.Version)
}

type IFile interface {
	AddFile(context.Context, FileReceipt) error
	GetFile(context.Context, string, Options) (FileReceipt, error)
	GetPiece(context.Context, string, Options) (PieceReceipt, error)
	GetReplica(context.Context, string, io.Writer, Options) (ReplicaReceipt, error)

	ListFile(context.Context, common.Address, Options) ([]FileReceipt, error)
	ListPiece(context.Context, common.Address, Options) ([]PieceReceipt, error)
	ListReplica(context.Context, common.Address, Options) ([]ReplicaReceipt, error)
}

type IPieceStore interface {
	PutPiece(context.Context, PieceCore, []byte, bool) error
	GetPiece(context.Context, string, io.Writer, Options) (PieceReceipt, error)
	GetPieceBySerial(context.Context, uint64) (string, error)

	PutReplica(context.Context, ReplicaCore, []byte, bool) error
	GetReplica(context.Context, string, io.Writer, Options) (ReplicaCore, error)
	GetReplicaBySerial(context.Context, uint64) (string, error)

	DeleteData(ctx context.Context, name string) error

	PutReplicaWitness(context.Context, common.Address, ReplicaWitness) error
	GetReplicaWitness(context.Context, common.Address, uint64) (ReplicaWitness, error)

	PutFile(context.Context, FileCore, []string, bool) error
	GetFile(context.Context, string, Options) (FileReceipt, error)
}

type ICommitment interface {
	Add(ICommitment) error
	Serialize() []byte // 48
	Raw() []byte       // 96
}

type IProof interface {
	Type() int
	Add(IProof) error
	Serialize() []byte // 48+48
}

type IChallenge interface {
	Type() int
	Add(ICommitment) error
	Commitment() ICommitment
}

type IPublicKey interface {
	VerifyKey() IVerifyKey

	GenCommitments(int, io.Reader) ([]ICommitment, error)
	GenCommitment(int, []byte, int) (ICommitment, error)
	GenProofs(IChallenge, int, io.Reader) ([]IProof, error)
	GenProof(IChallenge, int, []byte) (IProof, error)

	Serialize() []byte
	Deserialize([]byte) error
}

type IVerifyKey interface {
	VerifyProof(IChallenge, IProof) error

	Serialize() []byte
	Deserialize([]byte) error
}

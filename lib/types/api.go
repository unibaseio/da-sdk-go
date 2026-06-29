package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

type Options struct {
	UserDefined map[string]string
}

type Auth struct {
	Type string
	Addr common.Address
	Time int64
	Hash []byte
	Sign []byte
	// Msg, when set, carries a human-readable EIP-4361 / SIWE message that was
	// personal_signed verbatim (instead of the legacy Hash||be64(Time) bytes).
	// The verifier recovers the signer over Msg and reads the timestamp from it.
	// Empty for legacy clients — kept omitempty so their envelopes are unchanged.
	Msg string `json:",omitempty"`
}

type ModelResult struct {
	ModelMeta
}

type SpaceResult struct {
	SpaceMeta
}

type ListModelResult struct {
	Models []ModelResult
}

type ListSpaceResult struct {
	Spaces []SpaceResult
}

type ListGPUResult struct {
	GPUs []GPUMeta
}

type AccountResult struct {
	Name  common.Address
	Value *big.Int
}

type ChalResult struct {
	Epoch  int64
	Random [32]byte
}

type ListPieceResult struct {
	Pieces []PieceReceipt
}

type ListFileResult struct {
	Files []FileReceipt
}

type ListEdgeResult struct {
	Edges []EdgeReceipt
}

type ListReplicaResult struct {
	Replicas []ReplicaReceipt
}

type ListProofResult struct {
	Proofs []ProofStat
}

type ListStatResult struct {
	Stats []StatResult
}

type ListObjectResult struct {
	IsTruncated bool

	NextMarker string

	Objects []ObjectInfo

	Prefixes []string
}

type StatResult struct {
	Value int64
	Time  time.Time
}

func (hr *StatResult) Serialize() ([]byte, error) {
	return cbor.Marshal(hr)
}

func (hr *StatResult) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, hr)
}

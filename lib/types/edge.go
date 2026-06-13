package types

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

const (
	StoreType      = "store"
	StreamType     = "stream"
	ValidatorType  = "validator"
	ComputeType    = "compute"
	GatewayType    = "gateway"
	ClientType     = "client"
	HubType        = "hub"
	DownloaderType = "downloader"
)

type EdgeReceipt struct {
	EdgeMeta
	OnChain bool
	Revenue *big.Int
	Last    time.Time
}

// StoreStat is the per-node storage summary the gateway aggregates from its
// GormReplica mirror (what a store node currently holds on one chain).
type StoreStat struct {
	Store        common.Address
	ChainType    string
	ReplicaCount int64 // live (non-fake) replicas stored on this node
	FakeCount    int64 // replicas flagged fake (Forge/EPFake)
	PieceCount   int64 // distinct pieces this node contributes a replica to
	TotalSize    int64 // bytes of live replicas (sum of replica Size)
}

type EdgeMeta struct {
	ChainType string
	Type      string // "store", "compute"
	Name      common.Address
	PublicKey []byte
	ExposeURL string
	Hardware  HardwareInfo
}

func (em *EdgeMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(em)
}

func (em *EdgeMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, em)
}

type HardwareInfo struct {
	CPU    string
	Memory string
}

func (gm *HardwareInfo) Serialize() ([]byte, error) {
	return cbor.Marshal(gm)
}

func (gm *HardwareInfo) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, gm)
}

type IEdge interface {
	Register(context.Context, EdgeMeta) error
	GetEdge(context.Context, common.Address, Options) (EdgeReceipt, error)
	ListEdge(context.Context, Options) ([]EdgeReceipt, error)
}

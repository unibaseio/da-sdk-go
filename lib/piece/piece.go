package piece

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/common"
)

var logger = log.Logger("pieceStore")

var _ types.IPieceStore = &pieceStore{}

type pieceStore struct {
	ds types.IKVStore
	fs types.IFileStore
}

func New(ds types.IKVStore, fs types.IFileStore) *pieceStore {
	logger.Info("create piece store")

	ps := &pieceStore{
		ds: ds,
		fs: fs,
	}

	return ps
}

// DsPiece/PieceName; PieceCore
// DsPiece/PieceName/i; ReplicaName
// DsPiece/serial; PieceName
func (k *pieceStore) PutPiece(ctx context.Context, pc types.PieceCore, d []byte, update bool) error {
	logger.Debug("PutPiece: ", pc.Name)

	dsKey := types.NewKey(types.DsPiece, pc.Name)
	has, err := k.ds.Has(dsKey)
	if !update {
		if has && err == nil {
			return fmt.Errorf("already has piece: %s", pc.Name)
		}
	}

	pcb, err := pc.Serialize()
	if err != nil {
		return err
	}
	err = k.ds.Put(dsKey, pcb)
	if err != nil {
		return err
	}

	if pc.Serial > 0 {
		dsKey := types.NewKey(types.DsPiece, pc.Serial)
		pcnb, err := hex.DecodeString(pc.Name)
		if err != nil {
			return err
		}
		err = k.ds.Put(dsKey, pcnb)
		if err != nil {
			return err
		}
	}

	if len(d) > 0 {
		dskey := types.NewKey(types.DsFile, pc.Name)
		err = k.fs.Put(dskey, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k *pieceStore) GetPieceBySerial(ctx context.Context, serial uint64) (string, error) {
	dsKey := types.NewKey(types.DsPiece, serial)
	val, err := k.ds.Get(dsKey)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(val), nil
}

func (k *pieceStore) GetPiece(ctx context.Context, name string, w io.Writer, opt types.Options) (types.PieceReceipt, error) {
	logger.Debug("GetPiece: ", name, opt)
	res := types.PieceReceipt{}

	dsKey := types.NewKey(types.DsPiece, name)
	pcb, err := k.ds.Get(dsKey)
	if err != nil {
		return res, err
	}
	pc := new(types.PieceCore)
	err = pc.Deserialize(pcb)
	if err != nil {
		return res, err
	}
	res.PieceCore = *pc
	res.StoredOn = make([]common.Address, pc.Policy.N)
	res.Replicas = make([]string, pc.Policy.N)

	for i := uint8(0); i < pc.Policy.N; i++ {
		dsKey := types.NewKey(types.DsPiece, name, i)
		val, err := k.ds.Get(dsKey)
		if err != nil {
			continue
		}
		rname := hex.EncodeToString(val)
		res.Replicas[i] = rname

		dsKey = types.NewKey(types.DsReplica, rname)
		val, err = k.ds.Get(dsKey)
		if err != nil {
			continue
		}
		rc := new(types.ReplicaCore)
		err = rc.Deserialize(val)
		if err != nil {
			continue
		}
		res.StoredOn[i] = rc.StoredOn
	}

	if w != nil {
		dskey := types.NewKey(types.DsFile, name)
		val, err := k.fs.Get(dskey)
		if err != nil {
			return types.PieceReceipt{}, err
		}
		w.Write(val)
	}

	return res, nil
}

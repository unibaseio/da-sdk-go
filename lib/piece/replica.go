package piece

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"

	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/common"
)

// DsReplica/ReplicaName; ReplicaCore
// DsReplica/serial; ReplicaName
// DsReplica/addr/index; ReplicaWitness
func (k *pieceStore) PutReplica(ctx context.Context, rc types.ReplicaCore, d []byte, update bool) error {
	logger.Debug("PutReplica: ", rc)

	dsKey := types.NewKey(types.DsReplica, rc.Name)
	has, err := k.ds.Has(dsKey)
	if !update {
		if has && err == nil {
			return fmt.Errorf("already has replica: %s", rc.Name)
		}
	}

	rcb, err := rc.Serialize()
	if err != nil {
		return err
	}

	err = k.ds.Put(dsKey, rcb)
	if err != nil {
		return err
	}

	dsKey = types.NewKey(types.DsPiece, rc.Piece, rc.Index)
	rcnb, err := hex.DecodeString(rc.Name)
	if err != nil {
		return err
	}
	err = k.ds.Put(dsKey, rcnb)
	if err != nil {
		return err
	}

	if rc.Serial > 0 {
		dsKey := types.NewKey(types.DsReplica, rc.Serial)
		rcnb, err := hex.DecodeString(rc.Name)
		if err != nil {
			return err
		}
		err = k.ds.Put(dsKey, rcnb)
		if err != nil {
			return err
		}
	}

	// save data
	if len(d) > 0 {
		dskey := types.NewKey(types.DsFile, rc.Name)
		err = k.fs.Put(dskey, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k *pieceStore) GetReplicaBySerial(ctx context.Context, serial uint64) (string, error) {
	dsKey := types.NewKey(types.DsReplica, serial)
	val, err := k.ds.Get(dsKey)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(val), nil
}

func (k *pieceStore) GetReplica(ctx context.Context, name string, w io.Writer, opt types.Options) (types.ReplicaCore, error) {
	logger.Debug("GetReplica: ", name, opt)
	res := new(types.ReplicaCore)

	dsKey := types.NewKey(types.DsReplica, name)
	pcb, err := k.ds.Get(dsKey)
	if err != nil {
		return *res, err
	}
	err = res.Deserialize(pcb)
	if err != nil {
		return *res, err
	}

	if w != nil {
		// blob lives under DsFile (matches PutReplica/DeleteData + the piece-blob convention)
		dskey := types.NewKey(types.DsFile, name)
		data, err := k.fs.Get(dskey)
		if err != nil {
			return *res, err
		}

		if opt.UserDefined != nil {
			start := 0
			size := len(data)
			if opt.UserDefined["start"] != "" {
				start, _ = strconv.Atoi(opt.UserDefined["start"])
			}
			if opt.UserDefined["size"] != "" {
				size, _ = strconv.Atoi(opt.UserDefined["size"])
			}
			data = data[start : start+size]
		}
		w.Write(data)
	}

	return *res, nil
}

func (k *pieceStore) PutReplicaWitness(ctx context.Context, addr common.Address, rc types.ReplicaWitness) error {
	logger.Debug("PutReplicaWitness: ", addr, rc.Index)

	dsKey := types.NewKey(types.DsReplica, addr.String(), rc.Index)

	rcb, err := rc.Serialize()
	if err != nil {
		return err
	}

	err = k.ds.Put(dsKey, rcb)
	if err != nil {
		return err
	}

	return nil
}

func (k *pieceStore) GetReplicaWitness(ctx context.Context, addr common.Address, index uint64) (types.ReplicaWitness, error) {
	logger.Debug("GetReplicaWitness: ", addr, index)
	rw := new(types.ReplicaWitness)

	dsKey := types.NewKey(types.DsReplica, addr.String(), index)

	val, err := k.ds.Get(dsKey)
	if err != nil {
		return *rw, err
	}
	err = rw.Deserialize(val)
	if err != nil {
		return *rw, err
	}

	return *rw, nil
}

func (k *pieceStore) DeleteData(ctx context.Context, name string) error {
	logger.Debug("DeleteData: ", name)
	dskey := types.NewKey(types.DsFile, name)
	return k.fs.Delete(dskey)
}

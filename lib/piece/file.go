package piece

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// DsFile/FileName; FileCore
// DsFile/FileName/i; PieceName
func (k *pieceStore) PutFile(ctx context.Context, fc types.FileCore, pieces []string, append bool) error {
	logger.Debug("PutFile: ", fc.Name)

	nameByte := sha256.Sum256([]byte(fc.Name))
	name := hex.EncodeToString(nameByte[:])

	dsKey := types.NewKey(types.DsFile, name)
	has, err := k.ds.Has(dsKey)
	if !append {
		if has && err == nil {
			return fmt.Errorf("already has file: %s", fc.Name)
		}
	}

	pcb, err := fc.Serialize()
	if err != nil {
		return err
	}

	err = k.ds.Put(dsKey, pcb)
	if err != nil {
		return err
	}

	for i, pname := range pieces {
		dsKey := types.NewKey(types.DsFile, name, i)
		pcnb, err := hex.DecodeString(pname)
		if err != nil {
			return err
		}
		err = k.ds.Put(dsKey, pcnb)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k *pieceStore) GetFile(ctx context.Context, nameOrigin string, opt types.Options) (types.FileReceipt, error) {
	logger.Debug("GetFile: ", nameOrigin)

	nameByte := sha256.Sum256([]byte(nameOrigin))
	name := hex.EncodeToString(nameByte[:])

	res := types.FileReceipt{}

	dsKey := types.NewKey(types.DsFile, name)
	val, err := k.ds.Get(dsKey)
	if err != nil {
		return res, err
	}
	fc := new(types.FileCore)
	err = fc.Deserialize(val)
	if err != nil {
		return res, err
	}
	res.FileCore = *fc

	for i := 0; ; i++ {
		dsKey := types.NewKey(types.DsFile, name, i)
		val, err = k.ds.Get(dsKey)
		if err != nil {
			return res, nil
		}
		res.Pieces = append(res.Pieces, hex.EncodeToString(val))
	}
}

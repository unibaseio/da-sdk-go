package repo

import (
	"github.com/unibaseio/da-sdk-go/lib/config"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

type Repo interface {
	Config() *config.Config
	ReplaceConfig(cfg *config.Config) error

	Key() *key.Key

	MetaStore() types.IKVStore   // store meta
	DataStore() types.IFileStore // store data files

	Path() string

	Close() error

	// repo return the repo
	Repo() Repo
}

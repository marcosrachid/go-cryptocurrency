package db

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

type DBs struct {
	BlockIndex *leveldb.DB
	Chainstate *leveldb.DB
	Mempool    *leveldb.DB
}

var Instance *DBs = nil

func Start() error {
	Instance = new(DBs)
	if Instance.BlockIndex == nil {
		db, err := leveldb.OpenFile(os.Getenv("DB_PATH")+"/block-index.db", nil)
		if err != nil {
			return err
		}
		Instance.BlockIndex = db
	}
	if Instance.Chainstate == nil {
		db, err := leveldb.OpenFile(os.Getenv("DB_PATH")+"/chainstate.db", nil)
		if err != nil {
			return err
		}
		Instance.Chainstate = db
	}
	if Instance.Mempool == nil {
		db, err := leveldb.OpenFile(os.Getenv("DB_PATH")+"/mempool.db", nil)
		if err != nil {
			return err
		}
		Instance.Mempool = db
	}
	return nil
}

func Stop() {
	Instance.BlockIndex.Close()
	Instance.Chainstate.Close()
	Instance.Mempool.Close()
}

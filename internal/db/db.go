package db

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

type DBs struct {
	Block     *leveldb.DB
	UTXOState *leveldb.DB
	Mempool   *leveldb.DB
}

var Instance *DBs = nil

func Start() error {
	Instance = new(DBs)
	if Instance.Block == nil {
		db, err := leveldb.OpenFile(os.Getenv("DB_PATH")+"/block.db", nil)
		if err != nil {
			return err
		}
		Instance.Block = db
	}
	if Instance.UTXOState == nil {
		db, err := leveldb.OpenFile(os.Getenv("DB_PATH")+"/utxo.db", nil)
		if err != nil {
			return err
		}
		Instance.UTXOState = db
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
	Instance.Block.Close()
	Instance.UTXOState.Close()
	Instance.Mempool.Close()
}

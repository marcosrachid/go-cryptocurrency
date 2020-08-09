package mempool

import (
	"encoding/json"

	"github.com/marcosrachid/go-cryptocurrency/internal/db"
	"github.com/marcosrachid/go-cryptocurrency/internal/models"
	"github.com/marcosrachid/go-cryptocurrency/pkg/utils"

	"github.com/syndtr/goleveldb/leveldb/iterator"
)

func Get(transactionId string) (*models.SimpleTransaction, error) {
	response, err := db.Instance.Mempool.Get([]byte(transactionId), nil)
	if err != nil {
		return nil, err
	}
	decompressed, err := utils.Decompress(response)
	if err != nil {
		return nil, err
	}
	transaction := &models.SimpleTransaction{}
	json.Unmarshal(decompressed, transaction)
	return transaction, nil
}

func Delete(transactionId string) error {
	return db.Instance.Mempool.Delete([]byte(transactionId), nil)
}

func Put(transactionId string, transaction models.SimpleTransaction) error {
	data, err := json.Marshal(transaction)
	if err != nil {
		return err
	}
	compressed := utils.Compress(data)
	return db.Instance.Mempool.Put([]byte(transactionId), compressed, nil)
}

func Iterator() iterator.Iterator {
	return db.Instance.Mempool.NewIterator(nil, nil)
}

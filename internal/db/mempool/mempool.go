package mempool

import (
	"encoding/json"
	"go-cryptocurrency/internal/db"
	"go-cryptocurrency/internal/models"
)

func Get(key string) (*models.SimpleTransaction, error) {
	response, err := db.Instance.Mempool.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	transaction := &models.SimpleTransaction{}
	json.Unmarshal(response, transaction)
	return transaction, nil
}

func Delete(key string) error {
	return db.Instance.Mempool.Delete([]byte(key), nil)
}

func Put(key string, transaction models.SimpleTransaction) error {
	data, err := json.Marshal(transaction)
	if err != nil {
		return err
	}
	return db.Instance.Mempool.Put([]byte(key), []byte(data), nil)
}

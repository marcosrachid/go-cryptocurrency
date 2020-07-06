package utxo

import (
	"encoding/json"
	"go-cryptocurrency/internal/db"
	"go-cryptocurrency/internal/models"
)

func Get(key string) (*[]models.TransactionOutput, error) {
	response, err := db.Instance.UTXOState.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	utxo := &[]models.TransactionOutput{}
	json.Unmarshal(response, utxo)
	return utxo, nil
}

func Delete(key string) error {
	return db.Instance.UTXOState.Delete([]byte(key), nil)
}

func Put(key string, utxo []models.TransactionOutput) error {
	data, err := json.Marshal(utxo)
	if err != nil {
		return err
	}
	return db.Instance.UTXOState.Put([]byte(key), []byte(data), nil)
}

package utxo

import (
	"encoding/json"
	"go-cryptocurrency/internal/db"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/pkg/utils"
)

func Get(key string) (*[]models.TransactionOutput, error) {
	response, err := db.Instance.UTXOState.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	decompressed, err := utils.Decompress(response)
	if err != nil {
		return nil, err
	}
	utxo := &[]models.TransactionOutput{}
	json.Unmarshal(decompressed, utxo)
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
	compressed := utils.Compress(data)
	return db.Instance.UTXOState.Put([]byte(key), compressed, nil)
}

func Add(key string, utxo []models.TransactionOutput) {
	list, err := Get(key)
	if err != nil {
		l := make([]models.TransactionOutput, 0)
		list = &l
	}
	for _, u := range utxo {
		*list = append(*list, u)
	}
	// log.Println(list)
	Put(key, *list)
}

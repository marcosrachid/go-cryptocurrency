package utxo

import (
	"encoding/json"

	"github.com/marcosrachid/go-cryptocurrency/internal/db"
	"github.com/marcosrachid/go-cryptocurrency/internal/models"
	"github.com/marcosrachid/go-cryptocurrency/pkg/utils"
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

func Add(utxo map[string][]models.TransactionOutput) {
	for reciepient, output := range utxo {
		transactions, err := Get(reciepient)
		if err != nil {
			Put(reciepient, output)
		} else {
			Put(reciepient, append(output, *transactions...))
		}
	}
}

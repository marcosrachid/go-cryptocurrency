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

func Add(utxo []models.TransactionOutput) {
	m := make(map[string][]models.TransactionOutput)
	for i := 0; i < len(utxo); i++ {
		if val, ok := m[utxo[i].Reciepient]; ok {
			m[utxo[i].Reciepient] = append(val, utxo[i])
		} else {
			m[utxo[i].Reciepient] = []models.TransactionOutput{utxo[i]}
		}
	}
	for k, v := range m {
		Put(k, v)
	}
}

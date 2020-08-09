package services

import (
	"fmt"
	"strconv"

	"github.com/marcosrachid/go-cryptocurrency/internal/db/mempool"
	"github.com/marcosrachid/go-cryptocurrency/internal/db/utxo"
	"github.com/marcosrachid/go-cryptocurrency/internal/models"
)

func SendTransaction(arguments []string) (*models.SimpleTransaction, error) {
	data := ""
	if len(arguments) > 2 {
		data = arguments[2]
	}
	value, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		return nil, err
	}
	pubKey, err := GetPublicKey()
	if err != nil {
		return nil, err
	}

	inputs, balance, err := getInputsAndBalance(pubKey)
	if err != nil {
		return nil, err
	}
	if balance < value {
		return nil, fmt.Errorf("Not enough balance")
	}
	transaction, err := models.CreateSimpleTransaction(pubKey, arguments[0], inputs, value, data)
	if err != nil {
		return nil, err
	}
	mempool.Put(transaction.TransactionId, *transaction)
	return transaction, nil
}

func getInputsAndBalance(sender string) ([]models.TransactionInput, float64, error) {
	balanceOutputs, err := utxo.Get(sender)
	if err != nil {
		return nil, 0.0, err
	}
	balance := 0.0
	inputs := make([]models.TransactionInput, 0)
	for _, output := range *balanceOutputs {
		inputs = append(inputs, models.TransactionInput{
			TransactionOutputId:      output.Id,
			UnspentTransactionOutput: output,
		})
		balance += output.Value
	}
	return inputs, balance, nil
}

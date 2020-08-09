package models

import (
	"fmt"
	"time"

	"github.com/marcosrachid/go-cryptocurrency/pkg/utils"
)

type Transaction interface {
	GetTransactionId() string
	GetOutputs() []TransactionOutput
	calculateHash()
}

type RewardTransaction struct {
	TransactionId string            `json:"transaction_id"`
	Value         float64           `json:"value"`
	Timestamp     int64             `json:"timestamp"`
	Coinbase      string            `json:"coinbase"`
	Output        TransactionOutput `json:"output"`
}

type SimpleTransaction struct {
	TransactionId string              `json:"transaction_id"`
	Value         float64             `json:"value"`
	Timestamp     int64               `json:"timestamp"`
	Data          string              `json:"data"`
	Sender        string              `json:"sender"`
	FeeValue      float64             `json:"fee_value"`
	Signature     string              `json:"signature"`
	Inputs        []TransactionInput  `json:"inputs"`
	Outputs       []TransactionOutput `json:"outputs"`
}

type TransactionInput struct {
	TransactionOutputId      string            `json:"transaction_output_id"`
	UnspentTransactionOutput TransactionOutput `json:"unspent_transaction_output"`
}

type TransactionOutput struct {
	Id         string  `json:"id"`
	Reciepient string  `json:"reciepient"`
	Value      float64 `json:"value"`
	Timestamp  int64   `json:"timestamp"`
}

func (t RewardTransaction) GetTransactionId() string {
	return t.TransactionId
}

func (t RewardTransaction) GetOutputs() []TransactionOutput {
	return []TransactionOutput{t.Output}
}

func (t *RewardTransaction) calculateHash() {
	t.TransactionId = utils.ApplySha256(t.Coinbase + fmt.Sprintf("%v", t.Output) + fmt.Sprintf("%d", t.Timestamp))
}

func (t *RewardTransaction) calculateCoinbase(difficulty uint8, coinbase string) {
	t.Coinbase = utils.ApplySha256(fmt.Sprintf("%d", difficulty) + fmt.Sprintf("%d", t.Timestamp) + coinbase)
}

func CreateRewardTransaction(reciepient string, rewardValue float64, difficulty uint8, coinbase string) RewardTransaction {
	t := time.Now()
	transactionOutput := TransactionOutput{"", reciepient, rewardValue, t.UnixNano()}
	transactionOutput.calculateHash()
	transaction := RewardTransaction{"", rewardValue, t.UnixNano(), "", transactionOutput}
	transaction.calculateCoinbase(difficulty, coinbase)
	transaction.calculateHash()
	return transaction
}

func CreateSimpleTransaction(sender string, reciepient string, inputs []TransactionInput, value float64, data string) (*SimpleTransaction, error) {
	t := time.Now()
	reciepientOutput := TransactionOutput{"", reciepient, value, t.UnixNano()}
	reciepientOutput.calculateHash()
	balance := 0.0
	for _, input := range inputs {
		balance += input.UnspentTransactionOutput.Value
	}
	restOutput := TransactionOutput{"", sender, balance - value, t.UnixNano()}
	restOutput.calculateHash()
	outputs := []TransactionOutput{
		reciepientOutput,
		restOutput,
	}
	transaction := SimpleTransaction{
		TransactionId: "",
		Value:         value,
		Timestamp:     t.UnixNano(),
		Data:          data,
		Sender:        sender,
		FeeValue:      0.0,
		Signature:     "",
		Inputs:        inputs,
		Outputs:       outputs,
	}
	transaction.calculateHash()
	signature, err := utils.GenerateSignature(transaction.TransactionId)
	if err != nil {
		return nil, err
	}
	transaction.Signature = signature
	return &transaction, nil
}

func ApplyFees(transaction *SimpleTransaction) {

}

func (t SimpleTransaction) GetTransactionId() string {
	return t.TransactionId
}

func (t SimpleTransaction) GetOutputs() []TransactionOutput {
	return t.Outputs
}

func (t *SimpleTransaction) calculateHash() {
	t.TransactionId = utils.ApplySha256(t.Sender + fmt.Sprintf("%v", t.Inputs) + fmt.Sprintf("%v", t.Outputs) + fmt.Sprintf("%d", t.Timestamp))
}

func (to *TransactionOutput) calculateHash() {
	to.Id = utils.ApplySha256(to.Reciepient + fmt.Sprintf("%.16f", to.Value) + fmt.Sprintf("%d", to.Timestamp))
}

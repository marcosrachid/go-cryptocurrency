package models

import (
	"fmt"
	"go-cryptocurrency/pkg/utils"
	"time"
)

type Transaction interface {
	GetTransactionId() string
	calculateHash(difficulty uint8)
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
	TransactionOutputId       string            `json:"transaction_output_id"`
	UnspecntTransactionOutput TransactionOutput `json:"unspent_transaction_output"`
}

type TransactionOutput struct {
	Id                  string  `json:"id"`
	Reciepient          string  `json:"reciepient"`
	Value               float64 `json:"value"`
	ParentTransactionId string  `json:"parent_transaction_id"`
}

func (t RewardTransaction) GetTransactionId() string {
	return t.TransactionId
}

func (t *RewardTransaction) calculateHash(difficulty uint8) {
	t.TransactionId = utils.ApplySha256(t.Coinbase + fmt.Sprintf("%v", t.Output) + fmt.Sprintf("%d", difficulty) + fmt.Sprintf("%d", t.Timestamp))
	t.Output.ParentTransactionId = t.TransactionId
}

func (t *RewardTransaction) calculateCoinbase(difficulty uint8, coinbase string) {
	t.Coinbase = utils.ApplySha256(fmt.Sprintf("%d", difficulty) + fmt.Sprintf("%d", t.Timestamp) + coinbase)
}

func CreateRewardTransaction(reciepient string, rewardValue float64, difficulty uint8, coinbase string) RewardTransaction {
	transactionOutput := TransactionOutput{"", reciepient, rewardValue, ""}
	transactionOutput.calculateHash()
	t := time.Now()
	transaction := RewardTransaction{"", rewardValue, t.Unix(), "", transactionOutput}
	transaction.calculateCoinbase(difficulty, coinbase)
	transaction.calculateHash(difficulty)
	return transaction
}

func (t SimpleTransaction) GetTransactionId() string {
	return t.TransactionId
}

func (t *SimpleTransaction) calculateHash(difficulty uint8) {
	t.TransactionId = utils.ApplySha256(t.Sender + fmt.Sprintf("%v", t.Inputs) + fmt.Sprintf("%v", t.Outputs) + fmt.Sprintf("%d", difficulty) + fmt.Sprintf("%d", t.Timestamp))
}

func (to *TransactionOutput) calculateHash() {
	to.Id = utils.ApplySha256(to.Reciepient + fmt.Sprintf("%.16f", to.Value))
}

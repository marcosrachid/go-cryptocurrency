package models

type Transaction interface {
	GetTransactionId() string
}

type RewardTransaction struct {
	TransactionId string            `json:"transaction_id"`
	Value         float64           `json:"value"`
	Timestamp     int64             `json:"timestamp"`
	Coinbases     string            `json:"coinbase"`
	Difficulty    int64             `json:"difficulty"`
	output        TransactionOutput `json:"output"`
}

type SimpleTransaction struct {
	TransactionId string              `json:"transaction_id"`
	Value         float64             `json:"value"`
	Timestamp     int64               `json:"timestamp"`
	Data          string              `json:"data"`
	Sender        string              `json:"sender"`
	FeeValue      float64             `json:"fee_value"`
	Signature     []byte              `json:"signature"`
	inputs        []TransactionInput  `json:"inputs"`
	outputs       []TransactionOutput `json:"outputs"`
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

func (t *RewardTransaction) GetTransactionId() string {
	return t.TransactionId
}

func (t *SimpleTransaction) GetTransactionId() string {
	return t.TransactionId
}

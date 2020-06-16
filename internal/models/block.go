package models

import (
	"go-cryptocurrency/pkg/utils"
	"time"
)

type Block interface {
	GenerateNextBlock(transactions []Transaction) SimpleBlock
	CalculateHash() string
}

type GenesisBlock struct {
	Index     int64  `json:"index"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prev_hash"`
	Nonce     int64  `json:"nonce"`
}

type SimpleBlock struct {
	Index     int64         `json:"index"`
	Timestamp int64         `json:"timestamp"`
	Data      []Transaction `json:"data"`
	Hash      string        `json:"hash"`
	PrevHash  string        `json:"prev_hash"`
	Nonce     int64         `json:"nonce"`
}

// GENESIS METHODS
func (b GenesisBlock) GenerateNextBlock(transactions []Transaction) SimpleBlock {
	var newBlock SimpleBlock

	t := time.Now()

	newBlock.Index = b.Index + 1
	newBlock.Timestamp = t.Unix()
	newBlock.PrevHash = b.Hash
	newBlock.Data = transactions
	newBlock.Nonce = 0
	newBlock.Hash = newBlock.CalculateHash()

	return newBlock
}

func (b GenesisBlock) IsValid() bool {
	if len(b.Data) <= 0 {
		return false
	}
	return true
}

func (b GenesisBlock) CalculateHash() string {
	record := string(b.Index) + string(b.Nonce) + string(b.Timestamp) + string(b.Data) + b.PrevHash
	return utils.ApplySha256(record)
}

// SIMPLE BLOCK METHODS
func (b SimpleBlock) GenerateNextBlock(transactions []Transaction) SimpleBlock {
	var newBlock SimpleBlock

	t := time.Now()

	newBlock.Index = b.Index + 1
	newBlock.Timestamp = t.Unix()
	newBlock.PrevHash = b.Hash
	newBlock.Data = transactions
	newBlock.Nonce = 0
	newBlock.Hash = newBlock.CalculateHash()

	return newBlock
}

func (b SimpleBlock) IsValid(oldBlock SimpleBlock) bool {
	if oldBlock.Index+1 != b.Index {
		return false
	}
	if oldBlock.Hash != b.PrevHash {
		return false
	}
	if len(b.Data) <= 0 {
		return false
	}
	return true
}

func (b SimpleBlock) CalculateHash() string {
	record := string(b.Index) + string(b.Nonce) + string(b.Timestamp) + getMerkleRoot(b.Data) + b.PrevHash
	return utils.ApplySha256(record)
}

func getMerkleRoot(transactions []Transaction) string {
	count := len(transactions)
	var previousTreeLayer []string
	for _, transaction := range transactions {
		previousTreeLayer = append(previousTreeLayer, transaction.GetTransactionId())
	}
	treeLayer := previousTreeLayer
	for count > 1 {
		treeLayer = make([]string, 0)
		for i := 1; i < len(previousTreeLayer); i++ {
			treeLayer = append(treeLayer, utils.ApplySha256(previousTreeLayer[i-1]+previousTreeLayer[i]))
		}
		count = len(treeLayer)
		previousTreeLayer = treeLayer
	}
	var merkleRoot string = ""
	if len(treeLayer) == 1 {
		merkleRoot = treeLayer[0]
	}
	return merkleRoot
}

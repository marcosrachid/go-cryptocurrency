package models

import (
	"go-cryptocurrency/pkg/utils"
	"time"
)

type Block struct {
	Index     uint64        `json:"index"`
	Timestamp int64         `json:"timestamp"`
	Data      []Transaction `json:"data"`
	Hash      string        `json:"hash"`
	PrevHash  string        `json:"prev_hash"`
	Nonce     uint64        `json:"nonce"`
}

func (b Block) GenerateNextBlock(transactions []Transaction) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = b.Index + 1
	newBlock.Timestamp = t.Unix()
	newBlock.PrevHash = b.Hash
	newBlock.Data = transactions
	newBlock.Nonce = 0
	newBlock.Hash = newBlock.CalculateHash()

	return newBlock
}

func (b Block) IsValid(oldBlock Block) bool {
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

func (b Block) CalculateHash() string {
	record := string(b.Index) + string(b.Nonce) + string(b.Timestamp) + b.getMerkleRoot() + b.PrevHash
	return utils.ApplySha256(record)
}

func (b Block) getMerkleRoot() string {
	count := len(b.Data)
	var previousTreeLayer []string
	for _, transaction := range b.Data {
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

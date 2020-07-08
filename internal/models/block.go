package models

import (
	"encoding/json"
	"fmt"
	"go-cryptocurrency/pkg/utils"
	"strings"
	"time"
)

type BlockData []Transaction

type Block struct {
	Height            uint64    `json:"height"`
	ReceivedTimestamp int64     `json:"received_timestamp"`
	Timestamp         int64     `json:"timestamp"`
	Data              BlockData `json:"data"`
	Hash              string    `json:"hash"`
	PrevHash          string    `json:"prev_hash"`
	MerkleRoot        string    `json:"merkle_root"`
	Difficulty        uint8     `json:"difficulty"`
	RelayedBy         string    `json:"relayed_by"`
	Nonce             uint64    `json:"nonce"`
}

func (b Block) GenerateNextBlock(relayedBy string, difficulty uint8, transactions []Transaction) Block {
	var newBlock Block

	t := time.Now()

	if b.ReceivedTimestamp == 0 {
		newBlock.Height = 0
		newBlock.PrevHash = "0"
	} else {
		newBlock.Height = b.Height + 1
		newBlock.PrevHash = b.Hash
	}
	newBlock.ReceivedTimestamp = t.UnixNano()
	newBlock.Difficulty = difficulty
	newBlock.Data = transactions
	newBlock.RelayedBy = relayedBy
	newBlock.Nonce = 0
	newBlock.calculateMerkleRoot()
	newBlock.calculateHash()

	return newBlock
}

func (b Block) IsValid(oldBlock *Block, difficulty uint8) bool {
	if b.Height != 0 && oldBlock.Height+1 != b.Height {
		return false
	}
	if b.Height != 0 && oldBlock.Hash != b.PrevHash {
		return false
	}
	difficultyString := ""
	for len(difficultyString) < int(difficulty) {
		difficultyString += "0"
	}
	if runes := []rune(b.Hash); string(runes[0:difficulty]) != difficultyString {
		return false
	}
	return true
}

func (b *Block) Mine(difficulty uint8) {
	difficultyString := ""
	for len(difficultyString) < int(difficulty) {
		difficultyString += "0"
	}
	for runes := []rune(b.Hash); string(runes[0:difficulty]) != difficultyString; runes = []rune(b.Hash) {
		b.Nonce++
		b.Timestamp = time.Now().UnixNano()
		b.calculateHash()
	}
}

func (b *Block) calculateHash() {
	b.Hash = utils.ApplySha256(fmt.Sprintf("%d", b.Height) + fmt.Sprintf("%d", b.Nonce) + fmt.Sprintf("%d", b.Timestamp) + b.MerkleRoot + b.PrevHash)
}

func (b *Block) calculateMerkleRoot() {
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
	b.MerkleRoot = merkleRoot
}

func (bd *BlockData) UnmarshalJSON(data []byte) error {
	transactions := make([]json.RawMessage, 0)
	err := json.Unmarshal(data, &transactions)
	if err != nil {
		return err
	}
	results := make(BlockData, 0)
	for _, v := range transactions {
		var t Transaction
		if !strings.Contains(string(v), "coinbase") {
			t = &SimpleTransaction{}
		} else {
			t = &RewardTransaction{}
		}
		results = append(results, t)
		err := json.Unmarshal(v, t)
		if err != nil {
			return err
		}
	}
	*bd = results
	return nil
}

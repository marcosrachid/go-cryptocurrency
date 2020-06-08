package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func (b Block) GenerateNextBlock(BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = b.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = b.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func (b Block) IsValid(oldBlock Block) bool {
	if oldBlock.Index+1 != b.Index {
		return false
	}

	if oldBlock.Hash != b.PrevHash {
		return false
	}

	if calculateHash(b) != b.Hash {
		return false
	}

	return true
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

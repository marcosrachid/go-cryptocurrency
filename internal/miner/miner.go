package miner

import (
	"encoding/json"
	"fmt"
	"go-cryptocurrency/internal/db/block"
	"go-cryptocurrency/internal/db/mempool"
	"go-cryptocurrency/internal/db/utxo"
	"go-cryptocurrency/internal/global"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/services"
	"go-cryptocurrency/pkg/utils"
)

func MineBlocks() {
	for {
		if global.CURRENT_BLOCK.Height != global.NETWORK_HEIGHT {
			fmt.Println("Node unsinchronized")
			continue
		}
		transactionPool := make([]models.Transaction, 0)
		iter := mempool.Iterator()
		for iter.Next() {
			decompressed, err := utils.Decompress(iter.Value())
			if err != nil {
				fmt.Println("Mempool corrupted")
				continue
			}
			transaction := &models.SimpleTransaction{}
			json.Unmarshal(decompressed, transaction)
			transactionPool = append(transactionPool, transaction)
		}
		iter.Release()
		err := iter.Error()
		if err != nil {
			fmt.Println(err)
			continue
		}
		publicKey, err := services.GetPublicKey()
		if err != nil {
			fmt.Println(err)
			continue
		}
		var difficulty uint8
		if global.CURRENT_BLOCK.Height != 0 && global.CURRENT_BLOCK.Height%global.DIFFICULTY_ADJUSTMENT_BLOCK == 0 {
			difficulty = adjustDifficulty()
		} else {
			difficulty = global.CURRENT_BLOCK.Difficulty
		}
		var sequence uint64
		transactions := make([]models.Transaction, 0)
		for _, t := range transactionPool {
			t.CalculateHash(difficulty, sequence)
			transactions = append(transactions, t)
			transactionsBytes, _ := json.Marshal(transactions)
			sequence++
			if len(transactionsBytes) > int(global.BLOCK_SIZE) {
				break
			}
		}
		circulatingSupply := float64(global.CURRENT_BLOCK.Height) * global.REWARD
		if circulatingSupply != global.SUPPLY_LIMIT {
			var reward models.RewardTransaction
			if global.REWARD > global.SUPPLY_LIMIT-circulatingSupply {
				reward = models.CreateRewardTransaction(publicKey, global.SUPPLY_LIMIT-circulatingSupply, difficulty, global.COINBASE, sequence)
			} else {
				reward = models.CreateRewardTransaction(publicKey, global.REWARD, difficulty, global.COINBASE, sequence)
			}
			transactions = append(transactions, &reward)
		}
		if len(transactions) == 0 {
			fmt.Println("No transaction to mine")
			continue
		}
		newBlock := global.CURRENT_BLOCK.GenerateNextBlock(publicKey, difficulty, transactions)
		newBlock.Mine(difficulty)
		if !newBlock.IsValid(global.CURRENT_BLOCK, difficulty) {
			fmt.Println("Invalid block")
			continue
		}
		err = block.Put(newBlock)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var u []models.TransactionOutput
		for _, t := range newBlock.Data {
			u = append(u, t.GetOutputs()...)
		}
		utxo.Add(publicKey, u)
		global.CURRENT_BLOCK = &newBlock
		global.NETWORK_HEIGHT = newBlock.Height
		fmt.Println("Block mined: ", newBlock.Hash)
	}
}

func adjustDifficulty() uint8 {
	fmt.Println("Adjusting difficulty...")
	currentBlock := global.CURRENT_BLOCK
	var timestamps [global.DIFFICULTY_ADJUSTMENT_BLOCK]uint64
	for i := int(global.DIFFICULTY_ADJUSTMENT_BLOCK - 1); i >= 0; i-- {
		timestamps[i] = currentBlock.Timestamp
		currentBlock, _ = block.GetByHeight(currentBlock.Height - 1)
	}
	var differences [global.DIFFICULTY_ADJUSTMENT_BLOCK - 1]uint64
	for k, t := range timestamps {
		if k > 0 {
			differences[k-1] = t - timestamps[k-1]
		}
	}
	average := utils.AverageUint64(differences[:])
	top := float64(global.MINING_TIME_RATE) * (1.0 + global.MINING_TIME_RATE_ERROR)
	bottom := float64(global.MINING_TIME_RATE) * (1.0 - global.MINING_TIME_RATE_ERROR)
	if average > top && global.CURRENT_BLOCK.Difficulty > 0 {
		return global.CURRENT_BLOCK.Difficulty - 1
	}
	if average < bottom && global.CURRENT_BLOCK.Difficulty < global.MAX_DIFFICULTY {
		return global.CURRENT_BLOCK.Difficulty + 1
	}
	return global.CURRENT_BLOCK.Difficulty
}

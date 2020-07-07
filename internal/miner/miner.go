package miner

import (
	"encoding/json"
	"fmt"
	"go-cryptocurrency/internal/db/block"
	"go-cryptocurrency/internal/db/mempool"
	"go-cryptocurrency/internal/global"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/services"
	"time"
)

func MineBlocks() {
	for {
		time.Sleep(5 * time.Second)
		if global.CURRENT_BLOCK.Height != global.NETWORK_HEIGHT {
			fmt.Println("Node unsinchronized")
			continue
		}
		transactionPool := make([]models.Transaction, 0)
		iter := mempool.Iterator()
		for iter.Next() {
			transaction := &models.SimpleTransaction{}
			json.Unmarshal(iter.Value(), transaction)
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
		var difficulty uint64
		if global.CURRENT_BLOCK.Height%global.DIFFICULTY_ADJUSTMENT_BLOCK == 0 {
			difficulty = adjustDifficulty()
		} else {
			difficulty = global.CURRENT_BLOCK.Difficulty
		}
		transactions := make([]models.Transaction, 0)
		for _, t := range transactionPool {
			transactions = append(transactions, t)
			transactionsBytes, _ := json.Marshal(transactions)
			if len(transactionsBytes) > int(global.BLOCK_SIZE) {
				break
			}
		}
		newBlock := global.CURRENT_BLOCK.GenerateNextBlock(publicKey, difficulty, transactions)
		newBlock.Mine(difficulty)
		if !newBlock.IsValid(global.CURRENT_BLOCK, difficulty) {
			fmt.Println("Invalid block")
			continue
		}
		circulatingSupply := float64(global.CURRENT_BLOCK.Height) * global.REWARD
		if circulatingSupply != global.SUPPLY_LIMIT {
			var reward models.RewardTransaction
			if global.REWARD > global.SUPPLY_LIMIT-circulatingSupply {
				reward = models.CreateRewardTransaction(publicKey, global.SUPPLY_LIMIT-circulatingSupply, difficulty, global.COINBASE)
			} else {
				reward = models.CreateRewardTransaction(publicKey, global.REWARD, difficulty, global.COINBASE)
			}
			newBlock.Data = append(newBlock.Data, &reward)
		}
		block.Put(newBlock)
		global.CURRENT_BLOCK = &newBlock
		global.NETWORK_HEIGHT = newBlock.Height
		fmt.Println("Block mined: ", newBlock.Hash)
	}
}

func adjustDifficulty() uint64 {
	fmt.Println("Adjusting difficulty...")
	return 0
}

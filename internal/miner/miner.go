package miner

import (
	"fmt"
	"go-cryptocurrency/internal/db/block"
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
		// preencher transactions
		publicKey, err := services.GetPublicKey()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if global.CURRENT_BLOCK.Height%global.DIFFICULTY_ADJUSTMENT_BLOCK == 0 {
			adjustDifficulty()
		}
		var sequence uint64 = 0
		transactions := make([]models.Transaction, 0)
		for _, t := range transactionPool {
			sequence++
			transactions = append(transactions, t)
			// se tamanho chegar a um mb entaum break
		}
		newBlock := global.CURRENT_BLOCK.GenerateNextBlock(publicKey, global.DIFFICULTY, transactions)
		newBlock.Mine(global.DIFFICULTY)
		if !newBlock.IsValid(global.CURRENT_BLOCK, global.DIFFICULTY) {
			fmt.Println("Invalid block")
			continue
		}
		reward := models.CreateRewardTransaction(publicKey, global.REWARD, sequence+1, global.DIFFICULTY, global.COINBASE)
		newBlock.Data = append(newBlock.Data, &reward)
		block.Put(newBlock)
		global.CURRENT_BLOCK = newBlock
		global.NETWORK_HEIGHT = newBlock.Height
		fmt.Println("Block mined: ", newBlock.Hash)
	}
}

func adjustDifficulty() {

}

package main

import (
	"log"
	"os"

	"go-cryptocurrency/internal/db"
	"go-cryptocurrency/internal/global"
	"go-cryptocurrency/internal/miner"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/network"
	"go-cryptocurrency/internal/network/handler"
	"go-cryptocurrency/internal/services"

	"github.com/joho/godotenv"
)

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = services.WalletStart()
	if err != nil {
		panic(err)
	}

	err = db.Start()
	if err != nil {
		panic(err)
	}
	defer db.Stop()

	err = global.Load()
	if err != nil {
		panic(err)
	}

	// CLI TO BE REMOVED
	go network.SocketServer(os.Getenv("CLI_PORT"), handler.CliHandler)
	go network.SocketServer(os.Getenv("NETWORK_PORT"), handler.DispatcherHandler)
	// Wait a minute to communicate with the network
	// time.Sleep(60 * time.Second)
	miner.MineBlocks()
}

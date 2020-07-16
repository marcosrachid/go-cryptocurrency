package main

import (
	"log"
	"os"

	"github.com/marcosrachid/go-cryptocurrency/internal/db"
	"github.com/marcosrachid/go-cryptocurrency/internal/global"
	"github.com/marcosrachid/go-cryptocurrency/internal/miner"
	"github.com/marcosrachid/go-cryptocurrency/internal/network"
	"github.com/marcosrachid/go-cryptocurrency/internal/network/handler"
	"github.com/marcosrachid/go-cryptocurrency/internal/services"

	"github.com/joho/godotenv"
)

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

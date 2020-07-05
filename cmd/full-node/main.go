package main

import (
	"os"

	"go-cryptocurrency/internal/db"
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
		panic(err)
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

	go network.SocketServer(os.Getenv("CLI_PORT"), handler.CliHandler)
	network.SocketServer(os.Getenv("NETWORK_PORT"), handler.DispatcherHandler)
}

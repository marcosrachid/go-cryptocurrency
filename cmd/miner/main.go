package main

import (
	"log"
	"os"

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

	go network.SocketServer(os.Getenv("CLI_PORT"), handler.CliHandler)
	network.SocketServer(os.Getenv("NETWORK_PORT"), handler.DispatcherHandler)
}

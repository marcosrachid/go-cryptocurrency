package main

import (
	"os"

	"go-cryptocurrency/internal/handler"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/network"

	"github.com/joho/godotenv"
)

var Blockchain []models.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = handler.WalletStart()
	if err != nil {
		panic(err)
	}

	network.SocketServer(os.Getenv("NETWORK_PORT"))
}

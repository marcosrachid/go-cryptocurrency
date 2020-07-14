package main

import (
	"os"

	"go-cryptocurrency/internal/db"
	"go-cryptocurrency/internal/global"
	"go-cryptocurrency/internal/network"
	"go-cryptocurrency/internal/network/handler"
	"go-cryptocurrency/internal/services"

	"github.com/joho/godotenv"
)

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

	err = global.Load()
	if err != nil {
		panic(err)
	}

	go network.SocketServer(os.Getenv("CLI_PORT"), handler.CliHandler)
	network.SocketServer(os.Getenv("NETWORK_PORT"), handler.DispatcherHandler)
}

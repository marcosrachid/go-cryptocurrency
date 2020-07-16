package main

import (
	"os"

	"github.com/marcosrachid/go-cryptocurrency/internal/db"
	"github.com/marcosrachid/go-cryptocurrency/internal/global"
	"github.com/marcosrachid/go-cryptocurrency/internal/network"
	"github.com/marcosrachid/go-cryptocurrency/internal/network/handler"
	"github.com/marcosrachid/go-cryptocurrency/internal/services"

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

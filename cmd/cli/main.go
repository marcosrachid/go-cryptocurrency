package main

import (
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/network"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	if len(argsWithProg) <= 1 {
		network.SocketClient(os.Getenv("NODE_HOST"), os.Getenv("CLI_PORT"), models.CLIRequest{})
		return
	}
	command := argsWithoutProg[0]
	arguments := argsWithoutProg[1:]
	network.SocketClient(os.Getenv("NODE_HOST"), os.Getenv("CLI_PORT"), models.CLIRequest{command, arguments})
}

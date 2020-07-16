package main

import (
	"os"

	"github.com/marcosrachid/go-cryptocurrency/internal/models"
	"github.com/marcosrachid/go-cryptocurrency/internal/network"

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
		network.SocketClient(os.Getenv("CLI_HOST"), os.Getenv("CLI_PORT"), models.CLIRequest{})
		return
	}
	command := argsWithoutProg[0]
	arguments := argsWithoutProg[1:]
	network.SocketClient(os.Getenv("CLI_HOST"), os.Getenv("CLI_PORT"), models.CLIRequest{command, arguments})
}

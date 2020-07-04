package main

import (
	"fmt"
	"go-cryptocurrency/internal/handler"
	"os"
	"strings"

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
		showHelp()
		return
	}
	command := argsWithoutProg[0]
	arguments := argsWithoutProg[1:]
	switch {
	case strings.Compare(command, "wallet") == 0:
		keyCommands(arguments)
	case strings.Compare(command, "help") == 0:
		showHelp()
	default:
		fmt.Printf("Command \"%s\" not found\n", argsWithoutProg)
	}
}

func showHelp() {
	fmt.Println(
		`
Command usage:
	<command> [arguments]
The commands are:
	block 			show block arguments
	config			show config arguments
	node			show node arguments
	transaction		show transaction arguments
	wallet			show key arguments
	help			shows help
		`,
	)
}

func showKeyHelp() {
	fmt.Println(
		`
Command usage:
	wallet [arguments]
The arguments are:
	public-key		show current public-key
	private-key		show current private-key
	new			generate new wallet
		`,
	)
}

func keyCommands(arguments []string) {
	switch {
	case len(arguments) <= 0:
		showKeyHelp()
	case strings.Compare(arguments[0], "public-key") == 0:
		handler.PrintPublicKey()
	case strings.Compare(arguments[0], "private-key") == 0:
		handler.PrintKey()
	case strings.Compare(arguments[0], "new") == 0:
		handler.WalletGenerate()
	default:
		fmt.Printf("Command \"!wallet %s\" not found\n", arguments)
	}
}

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
	case strings.Compare(command, "block") == 0:
		blockCommands(arguments)
	case strings.Compare(command, "network") == 0:
		networkCommands(arguments)
	case strings.Compare(command, "transaction") == 0:
		transactionCommands(arguments)
	case strings.Compare(command, "wallet") == 0:
		keyCommands(arguments)
	case strings.Compare(command, "-h") == 0 || strings.Compare(command, "--help") == 0:
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
	network			show network arguments
	transaction		show transaction arguments
	wallet			show key arguments
	-h, --help		shows help
		`,
	)
}

func showBlockHelp() {
	fmt.Println(
		`
Command usage:
	block [arguments]
	block current
	block 1
	block 7b86d8a6216ffb3eaa359d9f7358437dcd23a3703a8ede4e28783a1446f6da7d
The arguments are:
	current or [0-9]+ or [a-z0-9]{64}			show current block or from specified height or from specified hash
	-e, --height current or [a-z0-9]{64} 		show current height or from specified hash
	-x, --hash current or [0-9]+				show current hash or from specified height
	-h, --help						shows help
		`,
	)
}

func showNetworkHelp() {
	fmt.Println(
		`
Command usage:
	network [arguments]
The arguments are:
	-v, --version			show network version
	-s, --supply-limit		show supply limit
	-m, --minimum-transaction	show minimum transaction value accepted above zero
	-c, --circulating-supply	show current circulating supply
	-d, --difficulty		show current difficulty
	-r, --reward			show current reward value
	-f, --fees			show current fees percentage
	-b, --block-size		show block maximum size
	-h, --help			shows help
		`,
	)
}

func showTransactionHelp() {
	fmt.Println(
		`
Command usage:
	transaction [reciepient public-key] [decimal value]
	transaction [reciepient public-key] [decimal value] [string data]
	transaction [reciepient public-key] [decimal value] [string data]
	transaction {[reciepient public-key 1], [reciepient public-key 2], ...} {[decimal value 1], [decimal value 2], ...}
	transaction {[reciepient public-key 1], [reciepient public-key 2], ...} {[decimal value 1], [decimal value 2], ...} {[string data 1], [string data 2], ...}
Examples:
	transaction "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==" 1
	transaction "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==" 1 "any string"
	transaction "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==","MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==" 1,10
	transaction "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==","MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==" 1,10 "string 1","string 2"
		`,
	)
}

func showKeyHelp() {
	fmt.Println(
		`
Command usage:
	wallet [arguments]
The arguments are:
	-p, --public-key		show current public-key
	-k, --private-key		show current private-key
	-n, --new			generate new wallet
	-h, --help			shows help
		`,
	)
}

func blockCommands(arguments []string) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		showBlockHelp()
	default:
		fmt.Printf("Command \"block %s\" not found\n", arguments)
	}
}

func networkCommands(arguments []string) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		showNetworkHelp()
	default:
		fmt.Printf("Command \"network %s\" not found\n", arguments)
	}
}

func transactionCommands(arguments []string) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		showTransactionHelp()
	default:
		fmt.Printf("Command \"transaction %s\" not found\n", arguments)
	}
}

func keyCommands(arguments []string) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		showKeyHelp()
	case strings.Compare(arguments[0], "-p") == 0 || strings.Compare(arguments[0], "--public-key") == 0:
		handler.PrintPublicKey()
	case strings.Compare(arguments[0], "-k") == 0 || strings.Compare(arguments[0], "--private-key") == 0:
		handler.PrintKey()
	case strings.Compare(arguments[0], "-n") == 0 || strings.Compare(arguments[0], "--new") == 0:
		handler.WalletGenerate()
	default:
		fmt.Printf("Command \"wallet %s\" not found\n", arguments)
	}
}

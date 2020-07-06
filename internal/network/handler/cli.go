package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-cryptocurrency/internal/global"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/services"
	"strings"
)

const (
	HELP = `
Command usage:
	<command> [arguments]
The commands are:
	block 			show block arguments
	network			show network arguments
	transaction		show transaction arguments
	wallet			show key arguments
	-h, --help		shows help
`
	BLOCK_HELP = `
Command usage:
	block [arguments]
	block
	block 1
	block 7b86d8a6216ffb3eaa359d9f7358437dcd23a3703a8ede4e28783a1446f6da7d
The arguments are:
	empty or [0-9]+ or [a-z0-9]{64}			show current block or from specified height or from specified hash
	-e, --height empty or [a-z0-9]{64} 		show current height or from specified hash
	-x, --hash empty or [0-9]+			show current hash or from specified height
	-h, --help					shows help
`
	NETWORK_HELP = `
Command usage:
	network [arguments]
The arguments are:
	-s, --supply-limit		show supply limit
	-m, --minimum-transaction	show minimum transaction value accepted above zero
	-c, --circulating-supply	show current circulating supply
	-d, --difficulty		show current difficulty
	-r, --reward			show current reward value
	-f, --fees			show current fees percentage
	-b, --block-size		show block maximum size
	-h, --help			shows help
`
	TRANSACTION_HELP = `
Command usage:
	transaction [reciepient public-key] [decimal value]
	transaction [reciepient public-key] [decimal value] [string data]
	transaction [reciepient public-key] [decimal value] [string data]
Examples:
	transaction "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==" 1
	transaction "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEP2egppaZKvyJ2r6+B2vEBBwSQP0B
yiVVhTpH5PYh6vGiq8QGcqJOvtW6vq3fUGEEJdyXXi77EMgFP7LrdEIhYw==" 1 "any string"
`
	KEY_HELP = `
Command usage:
	wallet [arguments]
The arguments are:
	-p, --public-key				show current public-key
	-k, --private-key				show current private-key
	-n, --new					generate new wallet
	-i, --import [private-key]			import a wallet
	-b, --balance empty or [public-key] 		show balance from current wallet or specified public-key
	-h, --help					shows help
`
)

func CliHandler(data string, writer *bufio.Writer) {
	request := &models.CLIRequest{}
	json.Unmarshal([]byte(strings.Replace(data, global.END, "", 1)), request)
	switch {
	case strings.Compare(request.Command, "block") == 0:
		blockCommands(request.Arguments, writer)
	case strings.Compare(request.Command, "network") == 0:
		networkCommands(request.Arguments, writer)
	case strings.Compare(request.Command, "transaction") == 0:
		transactionCommands(request.Arguments, writer)
	case strings.Compare(request.Command, "wallet") == 0:
		keyCommands(request.Arguments, writer)
	case len(request.Command) == 0 || strings.Compare(request.Command, "-h") == 0 || strings.Compare(request.Command, "--help") == 0:
		answer(HELP, writer)
	default:
		answer(fmt.Sprintf("Command \"%s\" not found\n", request.Command), writer)
	}
}

func blockCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) > 0 && (strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0):
		answer(BLOCK_HELP, writer)
	case len(arguments) > 0 && (strings.Compare(arguments[0], "-e") == 0 || strings.Compare(arguments[0], "--height") == 0):
		response, err := services.GetHeight(arguments[1:])
		answerWithError(response, err, writer)
	case len(arguments) > 0 && (strings.Compare(arguments[0], "-x") == 0 || strings.Compare(arguments[0], "--hash") == 0):
		response, err := services.GetHash(arguments[1:])
		answerWithError(response, err, writer)
	default:
		response, err := services.GetBlock(arguments)
		answerWithError(response, err, writer)
	}
}

func networkCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(NETWORK_HELP, writer)
	case strings.Compare(arguments[0], "-s") == 0 || strings.Compare(arguments[0], "--supply-limit") == 0:
		answer(fmt.Sprintf("%.16f", global.SUPPLY_LIMIT), writer)
	case strings.Compare(arguments[0], "-m") == 0 || strings.Compare(arguments[0], "--minimum-transaction") == 0:
		answer(fmt.Sprintf("%.16f", global.MINIMUM_TRANSACTION), writer)
	case strings.Compare(arguments[0], "-c") == 0 || strings.Compare(arguments[0], "--circulating-supply") == 0:
		answer(fmt.Sprintf("%.16f", global.CIRCULATING_SUPPLY), writer)
	case strings.Compare(arguments[0], "-d") == 0 || strings.Compare(arguments[0], "--difficulty") == 0:
		answer(fmt.Sprintf("%d", global.DIFFICULTY), writer)
	case strings.Compare(arguments[0], "-r") == 0 || strings.Compare(arguments[0], "--reward") == 0:
		answer(fmt.Sprintf("%.16f", global.REWARD), writer)
	case strings.Compare(arguments[0], "-f") == 0 || strings.Compare(arguments[0], "--fees") == 0:
		answer(fmt.Sprintf("%.16f%%", global.FEES*100), writer)
	case strings.Compare(arguments[0], "-b") == 0 || strings.Compare(arguments[0], "--block-size") == 0:
		answer(fmt.Sprintf("%d bytes", global.BLOCK_SIZE), writer)
	default:
		answer(fmt.Sprintf("Command \"block %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func transactionCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(TRANSACTION_HELP, writer)
	default:
		response, err := services.SendTransaction(arguments)
		answerWithError(response, err, writer)
	}
}

func keyCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(KEY_HELP, writer)
	case strings.Compare(arguments[0], "-p") == 0 || strings.Compare(arguments[0], "--public-key") == 0:
		response, err := services.GetPublicKey()
		answerWithError(response, err, writer)
	case strings.Compare(arguments[0], "-k") == 0 || strings.Compare(arguments[0], "--private-key") == 0:
		response, err := services.GetKey()
		answerWithError(response, err, writer)
	case strings.Compare(arguments[0], "-n") == 0 || strings.Compare(arguments[0], "--new") == 0:
		response, err := services.WalletGenerate()
		answerWithError(response, err, writer)
	case strings.Compare(arguments[0], "-i") == 0 || strings.Compare(arguments[0], "--import") == 0:
		response, err := services.WalletImport(arguments[1:])
		answerWithError(response, err, writer)
	case strings.Compare(arguments[0], "-b") == 0 || strings.Compare(arguments[0], "--balance") == 0:
		response, err := services.Balance(arguments[1:])
		answerWithError(response, err, writer)
	default:
		answer(fmt.Sprintf("Command \"wallet %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func answer(data string, writer *bufio.Writer) {
	writer.Write([]byte(data))
	writer.Write([]byte(global.END))
	writer.Flush()
}

func answerWithError(data string, err error, writer *bufio.Writer) {
	if err != nil {
		answer(fmt.Sprintf("%v", err), writer)
	} else {
		answer(data, writer)
	}
}

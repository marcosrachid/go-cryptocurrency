package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/marcosrachid/go-cryptocurrency/internal/global"
	"github.com/marcosrachid/go-cryptocurrency/internal/models"
	"github.com/marcosrachid/go-cryptocurrency/internal/services"
)

const (
	help = `
Command usage:
	<command> [arguments]
The commands are:
	block 			show block arguments
	network			show network arguments
	transaction		show transaction arguments
	wallet			show key arguments
	-h, --help		shows help
`
	blockHelp = `
Command usage:
	block [arguments]
The arguments are:
	?([0-9]+|[a-z0-9]{64})			show current block or from specified height or from specified hash
	-e, --height ?[a-z0-9]{64} 		show current height or from specified hash
	-x, --hash ?[0-9]+			show current hash or from specified height
	-h, --help				shows help
`
	networkHelp = `
Command usage:
	network [arguments]
The arguments are:
	-s, --supply-limit		show supply limit
	-c, --circulating-supply	show current circulating supply
	-d, --difficulty		show current difficulty
	-r, --reward			show current reward value
	-f, --fees			show current fees percentage
	-b, --block-size		show block maximum size
	-h, --help			shows help
`
	transactionHelp = `
Command usage:
	transaction [reciepient public-key] [decimal value] ?[arguments]
The arguments are:
	-p, --priority [1|2|3] 		set the priority of transaction(being 1 the highest priority)
	-d, --data (.)+			set generic data to be placed on transaction
`
	keyHelp = `
Command usage:
	wallet [arguments]
The arguments are:
	-p, --public-key				show current public-key
	-k, --private-key				show current private-key
	-n, --new					generate new wallet
	-i, --import [private-key hex string]		import a wallet
	-b, --balance ?[public-key] 			show balance from current wallet or specified public-key
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
		answer(help, writer)
	default:
		answer(fmt.Sprintf("Command \"%s\" not found\n", request.Command), writer)
	}
}

func blockCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) > 0 && (strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0):
		answer(blockHelp, writer)
	case len(arguments) > 0 && (strings.Compare(arguments[0], "-e") == 0 || strings.Compare(arguments[0], "--height") == 0):
		response, err := services.GetHeight(arguments[1:])
		answerWithError(fmt.Sprintf("%d", response), err, writer)
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
		answer(networkHelp, writer)
	case strings.Compare(arguments[0], "-s") == 0 || strings.Compare(arguments[0], "--supply-limit") == 0:
		answer(fmt.Sprintf(global.DECIMAL_STRING, global.SUPPLY_LIMIT), writer)
	case strings.Compare(arguments[0], "-c") == 0 || strings.Compare(arguments[0], "--circulating-supply") == 0:
		answer(fmt.Sprintf(global.DECIMAL_STRING, services.GetCirculatingSupply()), writer)
	case strings.Compare(arguments[0], "-d") == 0 || strings.Compare(arguments[0], "--difficulty") == 0:
		answer(fmt.Sprintf("%d", services.GetDifficulty()), writer)
	case strings.Compare(arguments[0], "-r") == 0 || strings.Compare(arguments[0], "--reward") == 0:
		answer(fmt.Sprintf(global.DECIMAL_STRING, global.REWARD), writer)
	case strings.Compare(arguments[0], "-f") == 0 || strings.Compare(arguments[0], "--fees") == 0:
		answer(fmt.Sprintf("P1: %.2f%%, P2: %.2f%%, P3: %.2f%%", global.P1_FEES*100, global.P2_FEES*100, global.P3_FEES*100), writer)
	case strings.Compare(arguments[0], "-b") == 0 || strings.Compare(arguments[0], "--block-size") == 0:
		answer(fmt.Sprintf("%d bytes", global.BLOCK_SIZE), writer)
	default:
		answer(fmt.Sprintf("Command \"block %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func transactionCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(transactionHelp, writer)
	default:
		response, err := services.SendTransaction(arguments)
		if err != nil {
			answerWithError("", err, writer)
		} else {
			r, _ := json.Marshal(response)
			answer(string(r), writer)
		}
	}
}

func keyCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(keyHelp, writer)
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
		answerWithError(fmt.Sprintf(global.DECIMAL_STRING, response), err, writer)
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

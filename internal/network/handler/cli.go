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
	block current
	block 1
	block 7b86d8a6216ffb3eaa359d9f7358437dcd23a3703a8ede4e28783a1446f6da7d
The arguments are:
	current or [0-9]+ or [a-z0-9]{64}			show current block or from specified height or from specified hash
	-e, --height current or [a-z0-9]{64} 		show current height or from specified hash
	-x, --hash current or [0-9]+				show current hash or from specified height
	-h, --help						shows help
`
	NETWORK_HELP = `
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
`
	TRANSACTION_HELP = `
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
`
	KEY_HELP = `
Command usage:
	wallet [arguments]
The arguments are:
	-p, --public-key		show current public-key
	-k, --private-key		show current private-key
	-n, --new			generate new wallet
	-h, --help			shows help
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
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(BLOCK_HELP, writer)
	default:
		answer(fmt.Sprintf("Command \"block %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func networkCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(NETWORK_HELP, writer)
	default:
		answer(fmt.Sprintf("Command \"block %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func transactionCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(TRANSACTION_HELP, writer)
	default:
		answer(fmt.Sprintf("Command \"transaction %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func keyCommands(arguments []string, writer *bufio.Writer) {
	switch {
	case len(arguments) <= 0 || strings.Compare(arguments[0], "-h") == 0 || strings.Compare(arguments[0], "--help") == 0:
		answer(KEY_HELP, writer)
	case strings.Compare(arguments[0], "-p") == 0 || strings.Compare(arguments[0], "--public-key") == 0:
		answer(services.GetPublicKey(), writer)
	case strings.Compare(arguments[0], "-k") == 0 || strings.Compare(arguments[0], "--private-key") == 0:
		answer(services.GetKey(), writer)
	case strings.Compare(arguments[0], "-n") == 0 || strings.Compare(arguments[0], "--new") == 0:
		answer(services.WalletGenerate(), writer)
	default:
		answer(fmt.Sprintf("Command \"wallet %s\" not found", strings.Join(arguments, " ")), writer)
	}
}

func answer(data string, writer *bufio.Writer) {
	writer.Write([]byte(data))
	writer.Flush()
}

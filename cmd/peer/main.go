package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"go-cryptocurrency/internal/handler"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/network"
	"go-cryptocurrency/pkg/utils"

	"github.com/gin-gonic/gin"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/joho/godotenv"
)

var Blockchain []models.Block

func showHelp() {
	fmt.Println(
		`
Command usage:
	<command> [arguments]
The commands are:
	!key		show key arguments
	!quit		to exit
	!help		shows help
		`,
	)
}

func showKeyHelp() {
	fmt.Println(
		`
Command usage:
	!key [arguments]
The arguments are:
	public-key		show current public-key
	private-key		show current private-key
	new				generate new wallet
		`,
	)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = handler.WalletStart()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	s := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Handle("GET", "/socket.io", network.NodeHandler(s))
	r.Handle("POST", "/socket.io", network.NodeHandler(s))
	r.Handle("WS", "/socket.io", network.NodeHandler(s))
	r.Handle("WSS", "/socket.io", network.NodeHandler(s))
	go r.Run(":" + os.Getenv("REST_PORT"))

	time.Sleep(1 * time.Second)
	showHelp()
Loop:
	for {
		fmt.Print(">> ")
		text := utils.ReadInput()
		switch {
		case strings.HasPrefix(text, "!key"):
			keyCommands(text)
		case strings.Compare(text, "!quit") == 0:
			break Loop
		case strings.Compare(text, "!help") == 0:
			showHelp()
		case strings.Compare(text, "") == 0:
		default:
			fmt.Printf("Command \"%s\" not found\n", text)
		}
	}
}

func keyCommands(text string) {
	r := regexp.MustCompile("[^\\s]+")
	splitedText := r.FindAllString(text, -1)
	switch {
	case len(splitedText) <= 1:
		showKeyHelp()
	case strings.Compare(splitedText[1], "public-key") == 0:
		publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
		if err != nil {
			panic(err)
		}
		fmt.Println(publicKeyString)
	case strings.Compare(splitedText[1], "private-key") == 0:
		privateKeyString, err := utils.GetKeyStringFromPEMKey()
		if err != nil {
			panic(err)
		}
		fmt.Println(privateKeyString)
	case strings.Compare(splitedText[1], "new") == 0:
		err := handler.WalletGenerate()
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("Command \"%s\" not found\n", text)
	}
}

package main

import (
	"fmt"
	"os"
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
	!public-key		show current public-key
	!private-key	show current private-key
	!quit			to exit
	!help			shows help
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
		case strings.Compare(text, "!public-key") == 0:
			publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
			if err != nil {
				panic(err)
			}
			fmt.Println(publicKeyString)
		case strings.HasPrefix(text, "!private-key"):
			privateKeyString, err := utils.GetKeyStringFromPEMKey()
			if err != nil {
				panic(err)
			}
			fmt.Println(privateKeyString)
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

package main

import (
	"os"

	"go-cryptocurrency/internal/handler"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/internal/network"

	"github.com/gin-gonic/gin"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/joho/godotenv"
)

var Blockchain []models.Block

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
	r.Run(":" + os.Getenv("REST_PORT"))
}

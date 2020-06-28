package network

import (
	"log"

	"github.com/gin-gonic/gin"
	gosocketio "github.com/graarh/golang-socketio"
)

func NodeHandler(socketServer *gosocketio.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		socketServer.On(gosocketio.OnConnection, node)
		socketServer.On(gosocketio.OnDisconnection, disconnection)
		socketServer.ServeHTTP(c.Writer, c.Request)
	}
}

func MinerHandler(socketServer *gosocketio.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		socketServer.On(gosocketio.OnConnection, miner)
		socketServer.On(gosocketio.OnDisconnection, disconnection)
		socketServer.ServeHTTP(c.Writer, c.Request)
	}
}

func node(socket *gosocketio.Channel) {
	log.Println("channel: ", socket)
	log.Println("Node Connected")
}

func miner(socket *gosocketio.Channel) {
	log.Println("channel: ", socket)
	log.Println("Node Connected")
}

func disconnection(socket *gosocketio.Channel) {
	log.Println("channel: ", socket)
	log.Println("Node Connected")
}

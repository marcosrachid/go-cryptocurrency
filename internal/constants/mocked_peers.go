package constants

import "go-cryptocurrency/internal/models"

var PEERS [1]models.Peer = [1]models.Peer{
	models.Peer{"main-node.coin.com", 8888},
}

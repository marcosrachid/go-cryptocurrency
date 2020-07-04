package constants

import "go-cryptocurrency/internal/models"

var PEERS []models.Peer = []models.Peer{
	models.Peer{"main-node.coin.com", 8888, "miner"},
}

var TRANSACTION_SEQUENCE = 0

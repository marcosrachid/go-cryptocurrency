package global

import "go-cryptocurrency/internal/models"

var PEERS []models.Peer = []models.Peer{
	models.Peer{"main-node.coin.com", 8888, "miner"},
}
var BANNED []models.Peer = make([]models.Peer, 0)
var TRANSACTION_SEQUENCE = 0

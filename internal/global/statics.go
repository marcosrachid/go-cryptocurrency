package global

import "go-cryptocurrency/internal/models"

var (
	// List of Peers to be connect
	PEERS []models.Peer = []models.Peer{
		models.Peer{"main-node-1.coin.com", 8888, "miner"},
	}
	// Current block registered on node
	CURRENT_BLOCK *models.Block = &models.Block{}
	// Public IP
	IP = ""
	// Current height received from network
	NETWORK_HEIGHT uint64 = 0
)

package global

import "github.com/marcosrachid/go-cryptocurrency/internal/models"

var (
	// PEERS List of Peers to be connect
	PEERS []models.Peer = []models.Peer{
		models.Peer{
			Domain: "main-node-1.coin.com",
			Port:   8888,
			Type:   "miner",
		},
	}
	// CURRENT_BLOCK Current block registered on node
	CURRENT_BLOCK *models.Block = &models.Block{}
	// IP Public IP
	IP = ""
	// NETWORK_HEIGHT Current height received from network
	NETWORK_HEIGHT uint64 = 0
)

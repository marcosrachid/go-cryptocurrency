package global

import "go-cryptocurrency/internal/models"

var (
	// List of Peers to be connect
	PEERS []models.Peer = []models.Peer{
		models.Peer{"main-node.coin.com", 8888, "miner"},
	}
	// Current circulating coin supply
	CIRCULATING_SUPPLY = 0.0
	// Current difficulty to mine
	DIFFICULTY = STARTING_DIFFICULTY
	// Current value reward
	REWARD = STARTING_REWARD
	// Current block registered on node
	CURRENT_BLOCK models.Block = models.Block{}
	// Current height received from network
	NETWORK_HEIGHT uint64 = 0
)

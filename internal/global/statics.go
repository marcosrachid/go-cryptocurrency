package global

import "go-cryptocurrency/internal/models"

var (
	PEERS []models.Peer = []models.Peer{
		models.Peer{"main-node.coin.com", 8888, "miner"},
	}
	BANNED                            = make([]models.Peer, 0)
	CIRCULATING_SUPPLY                = 0.0
	DIFFICULTY                        = STARTING_DIFFICULTY
	REWARD                            = STARTING_REWARD
	TRANSACTION_SEQUENCE              = 0
	CURRENT_BLOCK        models.Block = models.Block{}
)

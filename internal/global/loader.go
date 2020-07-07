package global

import (
	"go-cryptocurrency/internal/db/block"
)

func Load() {
	block, err := block.GetLast()
	if err != nil {
		return
	}
	CURRENT_BLOCK = block
	NETWORK_HEIGHT = block.Height
}

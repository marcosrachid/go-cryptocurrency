package services

import (
	"go-cryptocurrency/internal/global"
)

func GetCirculatingSupply() float64 {
	return float64(global.CURRENT_BLOCK.Height) * global.REWARD
}

func GetDifficulty() uint8 {
	return global.CURRENT_BLOCK.Difficulty
}

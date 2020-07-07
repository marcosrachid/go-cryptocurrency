package services

import (
	"go-cryptocurrency/internal/db/block"
	"go-cryptocurrency/internal/global"
)

func GetCirculatingSupply() (float64, error) {
	block, err := block.GetLast()
	return float64(block.Height) * global.REWARD, err
}

func GetDifficulty() (uint64, error) {
	block, err := block.GetLast()
	return block.Difficulty, err
}

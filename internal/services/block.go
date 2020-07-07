package services

import (
	"encoding/json"
	"go-cryptocurrency/internal/db/block"
	"go-cryptocurrency/internal/global"
	"regexp"
	"strconv"
)

func GetBlock(arguments []string) (string, error) {

	if len(arguments) <= 0 {
		data, _ := json.Marshal(global.CURRENT_BLOCK)
		return string(data), nil
	}
	if height, err := strconv.Atoi(arguments[0]); err == nil {
		block, err := block.GetByHeight(uint64(height))
		if err != nil {
			return "", err
		}
		data, _ := json.Marshal(block)
		return string(data), nil
	}
	if match, _ := regexp.MatchString("[a-z0-9]{64}", arguments[0]); match {
		block, err := block.GetByHash(arguments[0])
		if err != nil {
			return "", err
		}
		data, _ := json.Marshal(block)
		return string(data), nil
	}
	return "", nil
}

func GetHeight(arguments []string) (uint64, error) {
	if len(arguments) <= 0 {
		return global.CURRENT_BLOCK.Height, nil
	}
	if match, _ := regexp.MatchString("[a-z0-9]{64}", arguments[0]); match {
		block, err := block.GetByHash(arguments[0])
		if err != nil {
			return 0, err
		}
		return block.Height, err
	}
	return 0, nil
}

func GetHash(arguments []string) (string, error) {

	if len(arguments) <= 0 {
		return global.CURRENT_BLOCK.Hash, nil
	}
	if height, err := strconv.Atoi(arguments[0]); err == nil {
		block, err := block.GetByHeight(uint64(height))
		if err != nil {
			return "", err
		}
		return block.Hash, nil
	}
	return "", nil
}

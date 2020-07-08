package global

import (
	"go-cryptocurrency/internal/db/block"
	"go-cryptocurrency/pkg/utils"
)

func Load() error {
	block, err := block.GetLast()
	if err == nil {
		CURRENT_BLOCK = block
		NETWORK_HEIGHT = block.Height
	}
	ip, err := utils.GetPublicIP()
	if err != nil {
		return err
	}
	IP = ip
	return nil
}

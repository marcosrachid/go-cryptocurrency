package block

import (
	"encoding/json"
	"fmt"
	"go-cryptocurrency/internal/db"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/pkg/utils"
)

func GetByHeight(height uint64) (*models.Block, error) {
	sha256, err := db.Instance.Block.Get([]byte(fmt.Sprintf("%d", height)), nil)
	if err != nil {
		return nil, err
	}
	response, err := db.Instance.Block.Get(sha256, nil)
	if err != nil {
		return nil, err
	}
	block := &models.Block{}
	json.Unmarshal(response, block)
	return block, nil
}

func GetByHash(hash string) (*models.Block, error) {
	sha256, err := db.Instance.Block.Get([]byte(hash), nil)
	if err != nil {
		return nil, err
	}
	response, err := db.Instance.Block.Get(sha256, nil)
	if err != nil {
		return nil, err
	}
	block := &models.Block{}
	json.Unmarshal(response, block)
	return block, nil
}

func DeleteByHeight(height uint64) error {
	sha256, err := db.Instance.Block.Get([]byte(fmt.Sprintf("%d", height)), nil)
	if err != nil {
		return err
	}
	response, err := db.Instance.Block.Get(sha256, nil)
	if err != nil {
		return err
	}
	block := &models.Block{}
	json.Unmarshal(response, block)
	err = db.Instance.Block.Delete([]byte(fmt.Sprintf("%d", block.Height)), nil)
	if err != nil {
		return err
	}
	err = db.Instance.Block.Delete([]byte(block.Hash), nil)
	if err != nil {
		return err
	}
	return db.Instance.Block.Delete([]byte(sha256), nil)
}

func DeleteByHash(hash string) error {
	sha256, err := db.Instance.Block.Get([]byte(hash), nil)
	if err != nil {
		return err
	}
	response, err := db.Instance.Block.Get(sha256, nil)
	if err != nil {
		return err
	}
	block := &models.Block{}
	json.Unmarshal(response, block)
	err = db.Instance.Block.Delete([]byte(fmt.Sprintf("%d", block.Height)), nil)
	if err != nil {
		return err
	}
	err = db.Instance.Block.Delete([]byte(block.Hash), nil)
	if err != nil {
		return err
	}
	return db.Instance.Block.Delete([]byte(sha256), nil)
}

func Put(block models.Block) error {
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	sha256 := utils.ApplySha256(string(data))
	err = db.Instance.Block.Put([]byte(fmt.Sprintf("%d", block.Height)), []byte(sha256), nil)
	if err != nil {
		return err
	}
	err = db.Instance.Block.Put([]byte(block.Hash), []byte(sha256), nil)
	if err != nil {
		return err
	}
	return db.Instance.Block.Put([]byte(sha256), []byte(data), nil)
}

package services

import (
	"fmt"
	"go-cryptocurrency/internal/db/utxo"
	"go-cryptocurrency/internal/models"
	"go-cryptocurrency/pkg/utils"
	"log"
)

func WalletStart() error {
	key, err := utils.GetKeyFromPEMKey()
	if err != nil {
		log.Println("Wallet does not exist...\nAssigning a wallet")
		key, err = utils.GenerateKey()
		if err != nil {
			return err
		}
		publicKey := &key.PublicKey
		err = utils.SavePEMKey(key)
		if err != nil {
			return err
		}
		err = utils.SavePublicPEMKey(publicKey)
		if err != nil {
			return err
		}
		publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
		if err != nil {
			return err
		}
		log.Println("Wallet created...\n", publicKeyString)
	} else {
		publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
		if err != nil {
			return err
		}
		log.Println("Wallet found...\n", publicKeyString)
	}
	return err
}

func GetKey() (string, error) {
	privateKeyString, err := utils.GetKeyStringFromPEMKey()
	if err != nil {
		return "", err
	}
	return privateKeyString, nil
}

func GetPublicKey() (string, error) {
	publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
	if err != nil {
		return "", err
	}
	return publicKeyString, nil
}

func WalletGenerate() (string, error) {
	key, err := utils.GenerateKey()
	if err != nil {
		return "", err
	}
	publicKey := &key.PublicKey
	err = utils.SavePEMKey(key)
	if err != nil {
		return "", err
	}
	err = utils.SavePublicPEMKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
	if err != nil {
		return "", err
	}
	return publicKeyString, nil
}

func WalletImport(arguments []string) (string, error) {
	if len(arguments) <= 0 {
		return "", fmt.Errorf("Private key path is mandatory")
	}
	key, err := utils.ImportPEMKey(arguments[0])
	if err != nil {
		return "", err
	}
	publicKey := &key.PublicKey
	err = utils.SavePEMKey(key)
	if err != nil {
		return "", err
	}
	err = utils.SavePublicPEMKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
	if err != nil {
		return "", err
	}
	return publicKeyString, nil
}

func Balance(arguments []string) (float64, error) {
	var transactions []models.TransactionOutput = []models.TransactionOutput{}
	var balance float64 = 0.0
	if len(arguments) <= 0 {
		pubkey, err := GetPublicKey()
		if err != nil {
			return 0.0, err
		}
		t, err := utxo.Get(pubkey)
		if err != nil {
			return 0.0, nil
		}
		transactions = append(transactions, *t...)
	} else {
		t, err := utxo.Get(arguments[0])
		if err != nil {
			return 0.0, nil
		}
		transactions = append(transactions, *t...)
	}
	for _, t := range transactions {
		balance += t.Value
	}
	return balance, nil
}

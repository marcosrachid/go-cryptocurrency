package handler

import (
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

func WalletGenerate() error {
	key, err := utils.GenerateKey()
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
	return err
}

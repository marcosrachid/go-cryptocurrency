package handler

import (
	"fmt"
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

func PrintKey() {
	privateKeyString, err := utils.GetKeyStringFromPEMKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(privateKeyString)
}

func PrintPublicKey() {
	publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
	if err != nil {
		panic(err)
	}
	fmt.Println(publicKeyString)
}

func WalletGenerate() {
	key, err := utils.GenerateKey()
	if err != nil {
		panic(err)
	}
	publicKey := &key.PublicKey
	err = utils.SavePEMKey(key)
	if err != nil {
		panic(err)
	}
	err = utils.SavePublicPEMKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicKeyString, err := utils.GetPublicKeyStringFromPublicPEMKey()
	if err != nil {
		panic(err)
	}
	log.Println("Wallet created...\n", publicKeyString)
}

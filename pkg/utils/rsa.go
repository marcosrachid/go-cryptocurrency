package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	PATH             = "key"
	PRIVATE_PEM_NAME = "private.pem"
	PRIVATE_PEM_TYPE = "PRIVATE KEY"
	PUBLIC_PEM_NAME  = "public.pem"
	PUBLIC_PEM_TYPE  = "PUBLIC KEY"
)

func GenerateKey() (*rsa.PrivateKey, error) {
	reader := rand.Reader
	bitSize := 2048
	return rsa.GenerateKey(reader, bitSize)
}

func SavePEMKey(key *rsa.PrivateKey) error {
	outFile, err := os.Create(filepath.Join(PATH, filepath.Base(PRIVATE_PEM_NAME)))
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  PRIVATE_PEM_TYPE,
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	return err
}

func GetKeyFromPEMKey() (*rsa.PrivateKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(PATH, filepath.Base(PRIVATE_PEM_NAME)))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(pemString))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func SavePublicPEMKey(pubkey rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  PUBLIC_PEM_TYPE,
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(filepath.Join(PATH, filepath.Base(PUBLIC_PEM_NAME)))
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	return err
}

func GetPublicKeyFromPublicPEMKey() (*rsa.PublicKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(PATH, filepath.Base(PUBLIC_PEM_NAME)))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(pemString))
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetPublicKeyStringFromPublicPEMKey() (string, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(PATH, filepath.Base(PUBLIC_PEM_NAME)))
	if err != nil {
		return "", err
	}
	return string(pemString), nil
}

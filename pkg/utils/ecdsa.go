package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
)

type signature struct {
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

const (
	path           = "key"
	privatePemName = "private.pem"
	privatePemType = "PRIVATE KEY"
	publicPemName  = "public.pem"
	publicPemType  = "PUBLIC KEY"
)

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func SavePEMKey(key *ecdsa.PrivateKey) error {
	outFile, err := os.Create(filepath.Join(path, filepath.Base(privatePemName)))
	if err != nil {
		return err
	}
	defer outFile.Close()

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	var privateKey = &pem.Block{
		Type:  privatePemType,
		Bytes: keyBytes,
	}

	err = pem.Encode(outFile, privateKey)
	return err
}

func GetKeyFromPEMKey() (*ecdsa.PrivateKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(path, filepath.Base(privatePemName)))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(pemString))
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetKeyStringFromPEMKey() (string, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(path, filepath.Base(privatePemName)))
	if err != nil {
		return "", err
	}
	return string(pemString), nil
}

func SavePublicPEMKey(pubkey *ecdsa.PublicKey) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  publicPemType,
		Bytes: keyBytes,
	}

	pemfile, err := os.Create(filepath.Join(path, filepath.Base(publicPemName)))
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	return err
}

func GetPublicKeyFromPublicPEMKey() (*ecdsa.PublicKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(path, filepath.Base(publicPemName)))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(pemString))
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*ecdsa.PublicKey), nil
}

func GetPublicKeyStringFromPublicPEMKey() (string, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(path, filepath.Base(publicPemName)))
	if err != nil {
		return "", err
	}
	return string(pemString), nil
}

func GenerateSignature(transactionId string) (string, error) {
	key, err := GetKeyFromPEMKey()
	if err != nil {
		return "", err
	}
	transactionIdBytes := []byte(transactionId)
	hashed := sha256.Sum256(transactionIdBytes)
	r, s, err := ecdsa.Sign(rand.Reader, key, hashed[:])
	if err != nil {
		return "", err
	}

	signature := &signature{
		R: r,
		S: s,
	}

	signature_json, err := json.Marshal(signature)
	if err != nil {
		return "", err
	}

	signature_hex := hex.EncodeToString(signature_json)
	return signature_hex, nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, transactionId string, signature string) (bool, error) {
	transactionIdBytes := []byte(transactionId)
	hashed := sha256.Sum256(transactionIdBytes)
	var sig signature
	signature_json, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	json.Unmarshal(signature_json, &sig)
	return ecdsa.Verify(publicKey, hashed[:], sig.R, sig.S), nil
}

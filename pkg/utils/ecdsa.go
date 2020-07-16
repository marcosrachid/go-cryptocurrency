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

type Signature struct {
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

func ImportPEMKey(keyHexString string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := hex.DecodeString(keyHexString)
	if err != nil {
		return nil, err
	}
	key, err := x509.ParseECPrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetKeyFromPEMKey() (*ecdsa.PrivateKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(path, filepath.Base(privatePemName)))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemString)
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetKeyStringFromPEMKey() (string, error) {
	key, err := GetKeyFromPEMKey()
	if err != nil {
		return "", err
	}

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(keyBytes), nil
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
	block, _ := pem.Decode(pemString)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*ecdsa.PublicKey), nil
}

func GetPublicKeyStringFromPublicPEMKey() (string, error) {
	pubkey, err := GetPublicKeyFromPublicPEMKey()
	if err != nil {
		return "", err
	}

	keyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(keyBytes), nil
}

func GenerateSignature(transactionId string) (string, error) {
	key, err := GetKeyFromPEMKey()
	if err != nil {
		return "", err
	}
	transactionIDBytes := []byte(transactionId)
	hashed := sha256.Sum256(transactionIDBytes)
	r, s, err := ecdsa.Sign(rand.Reader, key, hashed[:])
	if err != nil {
		return "", err
	}

	signature := &Signature{
		R: r,
		S: s,
	}

	signatureJSON, err := json.Marshal(signature)
	if err != nil {
		return "", err
	}

	signatureHex := hex.EncodeToString(signatureJSON)
	return signatureHex, nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, transactionId string, signature string) (bool, error) {
	transactionIDBytes := []byte(transactionId)
	hashed := sha256.Sum256(transactionIDBytes)
	var sig Signature
	signatureJSON, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	json.Unmarshal(signatureJSON, &sig)
	return ecdsa.Verify(publicKey, hashed[:], sig.R, sig.S), nil
}

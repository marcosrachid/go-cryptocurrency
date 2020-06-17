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
	PATH             = "key"
	PRIVATE_PEM_NAME = "private.pem"
	PRIVATE_PEM_TYPE = "PRIVATE KEY"
	PUBLIC_PEM_NAME  = "public.pem"
	PUBLIC_PEM_TYPE  = "PUBLIC KEY"
)

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func SavePEMKey(key *ecdsa.PrivateKey) error {
	outFile, err := os.Create(filepath.Join(PATH, filepath.Base(PRIVATE_PEM_NAME)))
	if err != nil {
		return err
	}
	defer outFile.Close()

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	var privateKey = &pem.Block{
		Type:  PRIVATE_PEM_TYPE,
		Bytes: keyBytes,
	}

	err = pem.Encode(outFile, privateKey)
	return err
}

func GetKeyFromPEMKey() (*ecdsa.PrivateKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(PATH, filepath.Base(PRIVATE_PEM_NAME)))
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

func SavePublicPEMKey(pubkey *ecdsa.PublicKey) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  PUBLIC_PEM_TYPE,
		Bytes: keyBytes,
	}

	pemfile, err := os.Create(filepath.Join(PATH, filepath.Base(PUBLIC_PEM_NAME)))
	if err != nil {
		return err
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	return err
}

func GetPublicKeyFromPublicPEMKey() (*ecdsa.PublicKey, error) {
	pemString, err := ioutil.ReadFile(filepath.Join(PATH, filepath.Base(PUBLIC_PEM_NAME)))
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
	pemString, err := ioutil.ReadFile(filepath.Join(PATH, filepath.Base(PUBLIC_PEM_NAME)))
	if err != nil {
		return "", err
	}
	return string(pemString), nil
}

func GenerateSignature(input string) (string, error) {
	key, err := GetKeyFromPEMKey()
	if err != nil {
		return "", err
	}
	inputBytes := []byte(input)
	hashed := sha256.Sum256(inputBytes)
	r, s, err := ecdsa.Sign(rand.Reader, key, hashed[:])
	if err != nil {
		return "", err
	}

	signature := &Signature{
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

func VerifySignature(publicKey *ecdsa.PublicKey, input string, signature string) (bool, error) {
	inputBytes := []byte(input)
	hashed := sha256.Sum256(inputBytes)
	var sig Signature
	signature_json, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	json.Unmarshal(signature_json, &sig)
	return ecdsa.Verify(publicKey, hashed[:], sig.R, sig.S), nil
}

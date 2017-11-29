package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func Encrypt(pub *rsa.PublicKey, msg string) ([]byte, error) {
	rng := rand.Reader
	return rsa.EncryptPKCS1v15(rng, pub, []byte(msg))
}

func Decrypt(priv *rsa.PrivateKey, msg []byte) (string, error) {
	rng := rand.Reader
	bytes, err := rsa.DecryptPKCS1v15(rng, priv, msg)
	return string(bytes), err
}

func ReadPublicKey(path string) (*rsa.PublicKey, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(file)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, _ := publicKey.(*rsa.PublicKey)
	return key, nil
}

func ReadPrivateKey(path string) (*rsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(file)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := GenerateKeyPair()
	wallet := Wallet{*private, public}
	return &wallet
}

func GenerateKeyPair() (*ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// Combine X and Y coordinates to create a single byte slice public key
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return private, public
}

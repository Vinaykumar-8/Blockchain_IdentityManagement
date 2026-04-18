package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
)

const dbPath = "./tmp/badger"
const lastHashKey = "l"

type Blockchain struct {
	Tip []byte 
	Database  *badger.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(lastHashKey))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)
    fmt.Printf("Adding new block with data: %s\n", data)

	err = bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}
		err = txn.Set([]byte(lastHashKey), newBlock.Hash)
		bc.Tip = newBlock.Hash
		return err
	})
	if err != nil {
		log.Panic(err)
	}
}

func VerifySignature(publicKeyBytes []byte, data []byte, signature []byte) bool {
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	curve := elliptic.P256()
	x := new(big.Int).SetBytes(publicKeyBytes[:len(publicKeyBytes)/2])
	y := new(big.Int).SetBytes(publicKeyBytes[len(publicKeyBytes)/2:])
	rawPubKey := ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	hash := sha256.Sum256(data)
	return ecdsa.Verify(&rawPubKey, hash[:], r, s)
}

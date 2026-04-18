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
	Database *badger.DB
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
	if err != nil { log.Panic(err) }

	newBlock := NewBlock(data, lastHash)

	err = bc.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil { return err }
		err = txn.Set([]byte(lastHashKey), newBlock.Hash)
		bc.Tip = newBlock.Hash
		return err
	})
	if err != nil { log.Panic(err) }
}

func InitBlockchain() *Blockchain {
	var tip []byte
	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	if err != nil { log.Panic(err) }

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(lastHashKey))
		if err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found. Creating Genesis...")
			genesis := NewGenesisBlock()
			err = txn.Set(genesis.Hash, genesis.Serialize())
			err = txn.Set([]byte(lastHashKey), genesis.Hash)
			tip = genesis.Hash
		} else {
			err = item.Value(func(val []byte) error {
				tip = val
				return nil
			})
		}
		return err
	})
	if err != nil { log.Panic(err) }

	return &Blockchain{tip, db}
}

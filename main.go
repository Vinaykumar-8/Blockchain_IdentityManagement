package main

import "fmt"

func main() {
    bc := InitBlockchain()
    defer bc.Database.Close()

    wallet := NewWallet()
    fmt.Printf("Wallet Public Key: %x\n", wallet.PublicKey)
    
    bc.AddBlock("Initial Identity Record")
    fmt.Println("Done!")
}

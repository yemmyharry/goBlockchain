package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	Timestamp    int64
	Transactions []string
	PrevHash     string
	Nonce        int
}

func NewBlock(nounce int, prevHash string) *Block {
	block := &Block{
		Timestamp:    time.Now().UTC().UnixNano(),
		Transactions: []string{"testing"},
		PrevHash:     prevHash,
		Nonce:        nounce,
	}

	return block
}

func (b *Block) Print() {
	fmt.Printf("Timestamp: %d\n", b.Timestamp)
	fmt.Printf("Transactions: %s\n", b.Transactions)
	fmt.Printf("PrevHash: %s\n", b.PrevHash)
	fmt.Printf("Nonce: %d\n", b.Nonce)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	NewBlock(0, "first hash").Print()

	log.Println("My blocks")
	fmt.Println("Hello, World!")
}

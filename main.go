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

type Blockchain struct {
	TransactionPool []string
	Chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash string) *Block {
	block := NewBlock(nonce, prevHash)
	bc.Chain = append(bc.Chain, block)
	return block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.CreateBlock(100, "hash 1")
	return bc
}

func (bc Blockchain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("Chain %d \n", i)
		block.Print()
	}
}

func NewBlock(nonce int, prevHash string) *Block {
	block := &Block{
		Timestamp:    time.Now().UTC().UnixNano(),
		Transactions: []string{"testing"},
		PrevHash:     prevHash,
		Nonce:        nonce,
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
	blockchain := NewBlockchain()
	blockchain.CreateBlock(12, "hash 2")
	blockchain.Print()
	log.Println("My blocks")
}

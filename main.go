package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	Timestamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     [32]byte       `json:"prev_hash"`
	Nonce        int            `json:"nonce"`
}

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float32 `json:"amount"`
}

func NewTransaction(sender, recipient string, amount float32) *Transaction {
	return &Transaction{sender, recipient, amount}
}

func (tx *Transaction) Print() {
	fmt.Printf("%s \n", strings.Repeat("-", 30))
	fmt.Printf("%s sent %.1f to %s\n", tx.Sender, tx.Amount, tx.Recipient)
}

type Blockchain struct {
	TransactionPool []*Transaction
	Chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	block := NewBlock(nonce, prevHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, block)
	bc.TransactionPool = []*Transaction{}
	return block
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, amount float32) {
	tx := NewTransaction(sender, recipient, amount)
	bc.TransactionPool = append(bc.TransactionPool, tx)
}
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := &Blockchain{}
	bc.CreateBlock(100, b.Hash())
	return bc
}

func (bc Blockchain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("*", 25))
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	block := &Block{
		Timestamp:    time.Now().UTC().UnixNano(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        nonce,
	}

	return block
}

func (b *Block) Print() {
	fmt.Printf("Timestamp: %d\n", b.Timestamp)
	fmt.Printf("PrevHash: %x\n", b.PrevHash)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	for _, t := range b.Transactions {
		t.Print()
	}

}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.AddTransaction("Alice", "Bob", 100)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(12, previousHash)
	blockchain.Print()

	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(240, previousHash)
	blockchain.Print()
	log.Println("My blocks")
}

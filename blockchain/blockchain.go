package blockchain

import (
	"fmt"
	"goBlockchain/block"
	"goBlockchain/transactions"
	"strings"
	"time"
)

var MINING_DIFFICULTY = 3

type Blockchain struct {
	TransactionPool []*transactions.Transaction
	Chain           []*block.Block
}

func NewBlockchain() *Blockchain {
	b := block.Block{}
	bc := &Blockchain{}
	bc.CreateBlock(100, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *block.Block {
	block := block.NewBlock(nonce, prevHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, block)
	bc.TransactionPool = []*transactions.Transaction{}
	return block
}

func (bc *Blockchain) CopyTransactionPool() []*transactions.Transaction {
	var txs []*transactions.Transaction
	for _, tx := range bc.TransactionPool {
		txs = append(txs, transactions.NewTransaction(tx.Sender, tx.Recipient, tx.Amount))
	}
	return txs
}

func (bc *Blockchain) ValidProof(nonce int, prevHash [32]byte, transactions []*transactions.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guess := block.Block{
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        nonce,
	}
	guessHash := guess.Hash()
	guessHashStr := fmt.Sprintf("%x", guessHash)

	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactionPool := bc.CopyTransactionPool()
	lastBlock := bc.LastBlock()
	prevHash := lastBlock.Hash()
	nonce := 0
	for !bc.ValidProof(nonce, prevHash, transactionPool, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, amount float32) {
	tx := transactions.NewTransaction(sender, recipient, amount)
	bc.TransactionPool = append(bc.TransactionPool, tx)
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc Blockchain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("*", 25))
}

package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"goBlockchain/transactions"
	"time"
)

type Block struct {
	Timestamp    int64                       `json:"timestamp"`
	Transactions []*transactions.Transaction `json:"transactions"`
	PrevHash     [32]byte                    `json:"prev_hash"`
	Nonce        int                         `json:"nonce"`
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*transactions.Transaction) *Block {
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

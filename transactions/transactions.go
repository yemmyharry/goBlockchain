package transactions

import (
	"fmt"
	"strings"
)

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

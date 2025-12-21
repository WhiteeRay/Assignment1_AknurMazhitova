package bank

import (
	"fmt"
	"time"
)

type TransactionType string

const (
	TransactionDeposit    TransactionType = "Deposit"
	TransactionWithdrawal TransactionType = "Withdrawal"
	TransactionTransfer   TransactionType = "Transfer"
)

type Transaction struct {
	ID        uint64
	Type      TransactionType
	Amount    float64
	Balance   float64
	Timestamp time.Time
	Note      string
}

func NewTransaction(id uint64, transType TransactionType, amount, balance float64, note string) Transaction {
	return Transaction{
		ID:        id,
		Type:      transType,
		Amount:    amount,
		Balance:   balance,
		Timestamp: time.Now(),
		Note:      note,
	}
}

func (t Transaction) GetFormattedTimestamp() string {
	return t.Timestamp.Format("2006-01-02 15:04:05")
}

func (t Transaction) GetFormattedDetails() string {
	return fmt.Sprintf("[%s] %s: $%.2f | Balance: $%.2f | %s",
		t.GetFormattedTimestamp(),
		t.Type,
		t.Amount,
		t.Balance,
		t.Note)
}

func (t Transaction) IsDeposit() bool {
	return t.Type == TransactionDeposit
}

func (t Transaction) IsWithdrawal() bool {
	return t.Type == TransactionWithdrawal
}

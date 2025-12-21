package bank

import (
	"fmt"
)

type BankAccount struct {
	accountNumber string
	holderName    string
	balance       float64
	transactions  []Transaction
	nextTxnID     uint64
	isActive      bool
}

func NewBankAccount(accountNumber, holderName string, initialBalance float64) *BankAccount {
	account := &BankAccount{
		accountNumber: accountNumber,
		holderName:    holderName,
		balance:       0,
		transactions:  make([]Transaction, 0),
		nextTxnID:     1,
		isActive:      true,
	}

	if initialBalance > 0 {
		account.deposit(initialBalance, "Initial Deposit")
	}

	return account
}

func (b *BankAccount) Deposit(amount float64) error {
	if !b.isActive {
		return fmt.Errorf("account is not active")
	}

	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}

	b.deposit(amount, "Cash Deposit")
	return nil
}

func (b *BankAccount) deposit(amount float64, note string) {
	b.balance += amount
	txn := NewTransaction(b.nextTxnID, TransactionDeposit, amount, b.balance, note)
	b.transactions = append(b.transactions, txn)
	b.nextTxnID++
}

func (b *BankAccount) Withdraw(amount float64) error {
	if !b.isActive {
		return fmt.Errorf("account is not active")
	}

	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}

	if amount > b.balance {
		return fmt.Errorf("insufficient funds: balance $%.2f, requested $%.2f",
			b.balance, amount)
	}

	b.withdraw(amount, "Cash Withdrawal")
	return nil
}

func (b *BankAccount) withdraw(amount float64, note string) {
	b.balance -= amount
	txn := NewTransaction(b.nextTxnID, TransactionWithdrawal, amount, b.balance, note)
	b.transactions = append(b.transactions, txn)
	b.nextTxnID++
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) GetAccountNumber() string {
	return b.accountNumber
}

func (b *BankAccount) GetHolderName() string {
	return b.holderName
}

func (b *BankAccount) GetTransactions() []Transaction {
	return b.transactions
}

func (b *BankAccount) GetTransactionCount() int {
	return len(b.transactions)
}

func (b *BankAccount) GetLastTransaction() (Transaction, error) {
	if len(b.transactions) == 0 {
		return Transaction{}, fmt.Errorf("no transactions found")
	}
	return b.transactions[len(b.transactions)-1], nil
}

func (b *BankAccount) GetTotalDeposited() float64 {
	total := 0.0
	for _, txn := range b.transactions {
		if txn.IsDeposit() {
			total += txn.Amount
		}
	}
	return total
}

func (b *BankAccount) GetTotalWithdrawn() float64 {
	total := 0.0
	for _, txn := range b.transactions {
		if txn.IsWithdrawal() {
			total += txn.Amount
		}
	}
	return total
}

func (b *BankAccount) IsActive() bool {
	return b.isActive
}

func (b *BankAccount) CloseAccount() error {
	if b.balance > 0 {
		return fmt.Errorf("cannot close account with positive balance")
	}
	b.isActive = false
	return nil
}

func (b *BankAccount) GetAccountInfo() string {
	status := "Active"
	if !b.isActive {
		status = "Inactive"
	}
	return fmt.Sprintf("Account: %s | Holder: %s | Status: %s | Balance: $%.2f",
		b.accountNumber, b.holderName, status, b.balance)
}

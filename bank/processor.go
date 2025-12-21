package bank

import (
	"fmt"
	"strings"
)

type TransactionRequest struct {
	Type   string
	Amount float64
}

type TransactionProcessor struct {
	account *BankAccount
}

func NewTransactionProcessor(account *BankAccount) *TransactionProcessor {
	return &TransactionProcessor{
		account: account,
	}
}

func (tp *TransactionProcessor) ProcessTransactions(requests []TransactionRequest) []error {
	errors := make([]error, 0)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("=== Processing Batch Transactions ===")
	fmt.Println(strings.Repeat("=", 60))

	for i, req := range requests {
		fmt.Printf("\nTransaction %d/%d: %s $%.2f\n", i+1, len(requests), req.Type, req.Amount)

		var err error
		switch req.Type {
		case "deposit":
			err = tp.account.Deposit(req.Amount)
		case "withdraw":
			err = tp.account.Withdraw(req.Amount)
		default:
			err = fmt.Errorf("unknown transaction type: %s", req.Type)
		}

		if err != nil {
			fmt.Printf("  ❌ Failed: %v\n", err)
			errors = append(errors, err)
		} else {
			fmt.Printf("  ✓ Success - New Balance: $%.2f\n", tp.account.GetBalance())
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("Batch Complete: %d successful, %d failed\n", len(requests)-len(errors), len(errors))
	fmt.Printf("Final Balance: $%.2f\n", tp.account.GetBalance())
	fmt.Println(strings.Repeat("=", 60))

	return errors
}

func (tp *TransactionProcessor) ValidateTransactions(requests []TransactionRequest) error {
	simulatedBalance := tp.account.GetBalance()

	for i, req := range requests {
		if req.Amount <= 0 {
			return fmt.Errorf("transaction %d: amount must be positive", i+1)
		}

		switch req.Type {
		case "deposit":
			simulatedBalance += req.Amount
		case "withdraw":
			simulatedBalance -= req.Amount
			if simulatedBalance < 0 {
				return fmt.Errorf("transaction %d: would result in negative balance", i+1)
			}
		default:
			return fmt.Errorf("transaction %d: unknown type '%s'", i+1, req.Type)
		}
	}

	return nil
}

func (tp *TransactionProcessor) GetAccount() *BankAccount {
	return tp.account
}

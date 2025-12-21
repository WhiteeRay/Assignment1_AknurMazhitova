package bank

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BankUI struct {
	account   *BankAccount
	processor *TransactionProcessor
	scanner   *bufio.Scanner
}

func NewBankUI(account *BankAccount) *BankUI {
	return &BankUI{
		account:   account,
		processor: NewTransactionProcessor(account),
		scanner:   bufio.NewScanner(os.Stdin),
	}
}

func (ui *BankUI) ShowMenu() {
	for {
		ui.displayMenu()

		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		ui.scanner.Scan()

		if !ui.handleMenuChoice(choice) {
			break
		}
	}
}

func (ui *BankUI) displayMenu() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("=== Bank Account Management System ===")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println(ui.account.GetAccountInfo())
	fmt.Printf("Total Transactions: %d | Deposited: $%.2f | Withdrawn: $%.2f\n",
		ui.account.GetTransactionCount(),
		ui.account.GetTotalDeposited(),
		ui.account.GetTotalWithdrawn())
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("1. Deposit Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Check Balance")
	fmt.Println("4. View Transaction History")
	fmt.Println("5. View Account Summary")
	fmt.Println("6. Process Batch Transactions")
	fmt.Println("7. Exit")
}

func (ui *BankUI) handleMenuChoice(choice int) bool {
	switch choice {
	case 1:
		ui.handleDeposit()
	case 2:
		ui.handleWithdraw()
	case 3:
		ui.handleCheckBalance()
	case 4:
		ui.handleTransactionHistory()
	case 5:
		ui.handleAccountSummary()
	case 6:
		ui.handleBatchTransactions()
	case 7:
		fmt.Println("Thank you for banking with us. Goodbye!")
		return false
	default:
		fmt.Println("Invalid option. Please try again.")
	}
	return true
}

func (ui *BankUI) handleDeposit() {
	fmt.Println("\n--- Deposit Money ---")

	var amount float64
	fmt.Print("Enter deposit amount: $")
	fmt.Scan(&amount)

	if err := ui.account.Deposit(amount); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Successfully deposited $%.2f\n", amount)
		fmt.Printf("New Balance: $%.2f\n", ui.account.GetBalance())
	}
}

func (ui *BankUI) handleWithdraw() {
	fmt.Println("\n--- Withdraw Money ---")

	var amount float64
	fmt.Print("Enter withdrawal amount: $")
	fmt.Scan(&amount)

	if err := ui.account.Withdraw(amount); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Successfully withdrew $%.2f\n", amount)
		fmt.Printf("New Balance: $%.2f\n", ui.account.GetBalance())
	}
}

func (ui *BankUI) handleCheckBalance() {
	fmt.Println("\n--- Account Balance ---")
	fmt.Printf("Current Balance: $%.2f\n", ui.account.GetBalance())
}

func (ui *BankUI) handleTransactionHistory() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("=== Transaction History ===")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println(ui.account.GetAccountInfo())
	fmt.Println(strings.Repeat("-", 70))

	transactions := ui.account.GetTransactions()

	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	for i, txn := range transactions {
		fmt.Printf("%d. %s\n", i+1, txn.GetFormattedDetails())
	}

	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("Total Transactions: %d\n", len(transactions))
}

func (ui *BankUI) handleAccountSummary() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("=== Account Summary ===")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println(ui.account.GetAccountInfo())
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Total Deposited: $%.2f\n", ui.account.GetTotalDeposited())
	fmt.Printf("Total Withdrawn: $%.2f\n", ui.account.GetTotalWithdrawn())
	fmt.Printf("Net Change: $%.2f\n",
		ui.account.GetTotalDeposited()-ui.account.GetTotalWithdrawn())
	fmt.Printf("Transaction Count: %d\n", ui.account.GetTransactionCount())

	if lastTxn, err := ui.account.GetLastTransaction(); err == nil {
		fmt.Printf("\nLast Transaction:\n")
		fmt.Printf("  %s\n", lastTxn.GetFormattedDetails())
	}
	fmt.Println()
}

func (ui *BankUI) handleBatchTransactions() {
	fmt.Println("\n--- Batch Transaction Processing ---")
	fmt.Print("How many transactions to process? ")

	var count int
	fmt.Scan(&count)

	if count <= 0 {
		fmt.Println("Invalid count")
		return
	}

	requests := make([]TransactionRequest, 0, count)

	for i := 0; i < count; i++ {
		fmt.Printf("\nTransaction %d/%d:\n", i+1, count)

		var txnType string
		fmt.Print("Type (deposit/withdraw): ")
		fmt.Scan(&txnType)

		var amount float64
		fmt.Print("Amount: $")
		fmt.Scan(&amount)

		requests = append(requests, TransactionRequest{
			Type:   strings.ToLower(txnType),
			Amount: amount,
		})
	}

	if err := ui.processor.ValidateTransactions(requests); err != nil {
		fmt.Printf("\nValidation failed: %v\n", err)
		fmt.Println("Transactions not processed.")
		return
	}

	ui.processor.ProcessTransactions(requests)
}

func (ui *BankUI) readLine() string {
	ui.scanner.Scan()
	return strings.TrimSpace(ui.scanner.Text())
}

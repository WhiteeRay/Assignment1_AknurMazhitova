package main

import (
	"fmt"
	"strings"

	"github.com/WhiteeRay/Assignment1/bank"
	"github.com/WhiteeRay/Assignment1/company"
	"github.com/WhiteeRay/Assignment1/library"
	"github.com/WhiteeRay/Assignment1/shapes"
)

func main() {
	displayWelcome()

	for {
		displayMainMenu()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if !handleMainMenuChoice(choice) {
			break
		}
	}

	displayGoodbye()
}

func displayWelcome() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("WELCOME TO ASSIGNMENT 1: GO BASICS - PACKAGES, METHODS & INTERFACES")
	fmt.Println(strings.Repeat("=", 70))
}

func displayMainMenu() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("MAIN MENU")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("1. Library Management System")
	fmt.Println("2. Shapes & Interfaces Demonstration")
	fmt.Println("3. Employee Management System")
	fmt.Println("4. Bank Account Simulation")
	fmt.Println("5. Exit Application")
	fmt.Println(strings.Repeat("-", 70))
}

func handleMainMenuChoice(choice int) bool {
	switch choice {
	case 1:
		runLibrarySystem()
	case 2:
		runShapesDemo()
	case 3:
		runCompanySystem()
	case 4:
		runBankSystem()
	case 5:
		return false
	default:
		fmt.Println("Invalid choice. Please select 1-5.")
	}
	return true
}

func runLibrarySystem() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("STARTING LIBRARY MANAGEMENT SYSTEM")
	fmt.Println(strings.Repeat("=", 70))

	library := library.NewLibrary("Central Public Library")

	library.AddBook("B001", "The-Go-Programming-Language", "Alan-Donovan")
	library.AddBook("B002", "Clean-Code", "Robert-Martin")
	library.AddBook("B003", "Design-Patterns", "Gang-of-Four")
	library.AddBook("B004", "Introduction-to-Algorithms", "Thomas-Cormen")
	library.AddBook("B005", "The-Pragmatic-Programmer", "Andy-Hunt")

	fmt.Println("Library initialized with 5 sample books")

	ui := library.NewLibraryUI(library)
	ui.ShowMenu()
}

func runShapesDemo() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("STARTING SHAPES & INTERFACES DEMONSTRATION")
	fmt.Println(strings.Repeat("=", 70))

	shapes.DemonstrateShapes()

	fmt.Println("\nPress Enter to return to main menu...")
	fmt.Scanln()
	fmt.Scanln()
}

func runCompanySystem() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("STARTING EMPLOYEE MANAGEMENT SYSTEM")
	fmt.Println(strings.Repeat("=", 70))

	company := company.NewCompany("TechCorp International")

	fullTime1 := company.NewFullTimeEmployee(1001, "Alice-Johnson", "Engineering", 5500.00)
	company.AddEmployee(fullTime1)

	fullTime2 := company.NewFullTimeEmployee(1002, "Bob-Smith", "Marketing", 4800.00)
	company.AddEmployee(fullTime2)

	fullTime3 := company.NewFullTimeEmployee(1003, "Carol-Davis", "Finance", 5200.00)
	company.AddEmployee(fullTime3)

	partTime1 := company.NewPartTimeEmployee(2001, "David-Wilson", 28.00, 20.0)
	company.AddEmployee(partTime1)

	partTime2 := company.NewPartTimeEmployee(2002, "Emma-Brown", 25.00, 15.0)
	company.AddEmployee(partTime2)

	fmt.Println("Company initialized with 3 full-time and 2 part-time employees")

	ui := company.NewCompanyUI(company)
	ui.ShowMenu()
}

func runBankSystem() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("STARTING BANK ACCOUNT SIMULATION")
	fmt.Println(strings.Repeat("=", 70))

	account := bank.NewBankAccount("ACC-2024-001", "John-Doe", 1000.00)

	fmt.Println("Bank account created with initial deposit of $1000.00")

	processor := bank.NewTransactionProcessor(account)

	sampleTransactions := []bank.TransactionRequest{
		{Type: "deposit", Amount: 500.00},
		{Type: "withdraw", Amount: 200.00},
		{Type: "deposit", Amount: 150.00},
	}

	processor.ProcessTransactions(sampleTransactions)

	ui := bank.NewBankUI(account)
	ui.ShowMenu()
}

func displayGoodbye() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("Thank you for using the Assignment 1 Application!")
	fmt.Println("Goodbye!")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()
}

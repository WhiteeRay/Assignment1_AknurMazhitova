package main

import (
	"fmt"
	"strings"

	"github.com/WhiteeRay/Assignment1/bank"
	"github.com/WhiteeRay/Assignment1/company"
	"github.com/WhiteeRay/Assignment1/library"
	"github.com/WhiteeRay/Assignment1/parking2"
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
	fmt.Println("5. Car ParkingLot")
	fmt.Println("6. Exit Application")
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
		runCarSystem2()
	case 6:
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

	lib := library.NewLibrary("Astana Central Library")

	lib.AddBook("B001", "Go Programming Language", "Alan Donovan")
	lib.AddBook("B002", "Clean Code", "Robert Martin")
	lib.AddBook("B003", "Design Patterns", "Gang of Four")
	lib.AddBook("B004", "Introduction to Algorithms", "Thomas Cormen")
	lib.AddBook("B005", "The Pragmatic Programmer", "Andy Hunt")

	fmt.Println("Library initialized with 5 sample books")

	ui := library.NewLibraryUI(lib)
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

	comp := company.NewCompany("TechCorp Kazakhstan")

	fullTime1 := company.NewFullTimeEmployee(1001, "Aisha Zhumabekova", "Engineering", 550000.0)
	comp.AddEmployee(fullTime1)

	fullTime2 := company.NewFullTimeEmployee(1002, "Bek Nursultan", "Marketing", 480000.0)
	comp.AddEmployee(fullTime2)

	fullTime3 := company.NewFullTimeEmployee(1003, "Karol Dosanov", "Finance", 520000.0)
	comp.AddEmployee(fullTime3)

	partTime1 := company.NewPartTimeEmployee(2001, "David Ermekov", 28000.0, 20.0)
	comp.AddEmployee(partTime1)

	partTime2 := company.NewPartTimeEmployee(2002, "Emma Bekzhanova", 25000.0, 15.0)
	comp.AddEmployee(partTime2)

	fmt.Println("Company initialized with 3 full-time and 2 part-time employees")

	ui := company.NewCompanyUI(comp)
	ui.ShowMenu()
}

func runBankSystem() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("STARTING BANK ACCOUNT SIMULATION")
	fmt.Println(strings.Repeat("=", 70))

	account := bank.NewBankAccount("ACC-2025-001", "Zhanar Nurmaganbetova", 1000000.0)

	fmt.Println("Bank account created with initial deposit of 1,000,000₸")

	processor := bank.NewTransactionProcessor(account)

	sampleTransactions := []bank.TransactionRequest{
		{Type: "deposit", Amount: 500000.0},
		{Type: "withdraw", Amount: 200000.0},
		{Type: "deposit", Amount: 150000.0},
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




func runCarSystem2(){
	lot:= ParkingLot.NewParkingLot()
	rate:=500
	for{
		fmt.Println("1. Park Car")
		fmt.Println("2. UnPark Car")
		fmt.Println("3. Update Car")
		fmt.Println("4. List Cars")
		fmt.Println("5. Show Bills")
		fmt.Println("6. Exit")


		var choice int
		fmt.Scan(&choice)

		switch choice{
		case 1:
			var plate, owner string
			var hours int

			fmt.Print("Plate")
			fmt.Scan(&plate)
			fmt.Print("Owner")
			fmt.Scan(&owner)
			fmt.Print("Hours: ")
			fmt.Scan(&hours)

			err:= lot.ParkCar(ParkingLot.Car{
				Plate: plate,
				Owner: owner,
				HoursParked: hours,
			})
			if err!= nil{
				fmt.Println("Errors:",err)
			}

		case 2:
			var plate string
			fmt.Println("Plate: ")
			fmt.Scan(&plate)
			car, err :=lot.UnparkCar(plate)
			if err != nil{
				fmt.Println("Error: ", err)
			} else{
				fmt.Println("Unparked:", car.Plate)
			}
		case 3:
			var plate string
			var hours int
			fmt.Print("Plate: ")
			fmt.Scan(&plate)
			fmt.Print("Hours: ")
			fmt.Scan(&hours)
			err := lot.UpdateHours(plate,hours)
			if err!=nil{
				fmt.Println("Error: ", err)
			} else{
				fmt.Println("Updated: ", hours)
			}
		case 4:
			cars:=lot.ListCars()
			for _,car:=range cars{
				fmt.Printf("Plate: %s | Owner: %s | Hours: %d\n",
					car.Plate, car.Owner, car.HoursParked)
			}
		case 5:
			bills := lot.Bills(rate)
			for plate,cost:=range bills{
				fmt.Printf("%s -> %d\n", plate,cost)
			}
		case 6:
			return
		default:
			fmt.Println("Invalid Option")
		}



	}
	


}
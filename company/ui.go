package company

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CompanyUI struct {
	company *Company
	scanner *bufio.Scanner
}

func NewCompanyUI(company *Company) *CompanyUI {
	return &CompanyUI{
		company: company,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (ui *CompanyUI) ShowMenu() {
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

func (ui *CompanyUI) displayMenu() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("=== %s Employee Management System ===\n", ui.company.GetCompanyName())
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Total Employees: %d | Full-Time: %d | Part-Time: %d\n",
		ui.company.GetEmployeeCount(),
		ui.company.GetFullTimeCount(),
		ui.company.GetPartTimeCount())
	fmt.Printf("Monthly Payroll: $%.2f | Average Salary: $%.2f\n",
		ui.company.CalculateTotalPayroll(),
		ui.company.GetAverageSalary())
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("1. Add Full-Time Employee")
	fmt.Println("2. Add Part-Time Employee")
	fmt.Println("3. List All Employees")
	fmt.Println("4. List Full-Time Employees")
	fmt.Println("5. List Part-Time Employees")
	fmt.Println("6. Search Employee by ID")
	fmt.Println("7. Remove Employee")
	fmt.Println("8. Show Payroll Summary")
	fmt.Println("9. Exit")
}

func (ui *CompanyUI) handleMenuChoice(choice int) bool {
	switch choice {
	case 1:
		ui.handleAddFullTimeEmployee()
	case 2:
		ui.handleAddPartTimeEmployee()
	case 3:
		ui.handleListAllEmployees()
	case 4:
		ui.handleListFullTimeEmployees()
	case 5:
		ui.handleListPartTimeEmployees()
	case 6:
		ui.handleSearchEmployee()
	case 7:
		ui.handleRemoveEmployee()
	case 8:
		ui.handlePayrollSummary()
	case 9:
		fmt.Println("Exiting Employee Management System. Goodbye!")
		return false
	default:
		fmt.Println("Invalid option. Please try again.")
	}
	return true
}

func (ui *CompanyUI) handleAddFullTimeEmployee() {
	fmt.Println("\nAdd Full-Time Employee")

	var id uint64
	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)
	ui.scanner.Scan()

	fmt.Print("Enter Employee Name: ")
	name := ui.readLine()

	fmt.Print("Enter Department: ")
	department := ui.readLine()

	var salary float64
	fmt.Print("Enter Monthly Salary: $")
	fmt.Scan(&salary)

	employee := NewFullTimeEmployee(id, name, department, salary)

	if err := ui.company.AddEmployee(employee); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Full-Time Employee '%s' added successfully!\n", name)
	}
}

func (ui *CompanyUI) handleAddPartTimeEmployee() {
	fmt.Println("\nAdd Part-Time Employee")

	var id uint64
	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)
	ui.scanner.Scan()

	fmt.Print("Enter Employee Name: ")
	name := ui.readLine()

	var hourlyRate float64
	fmt.Print("Enter Hourly Rate: $")
	fmt.Scan(&hourlyRate)

	var hoursWorked float64
	fmt.Print("Enter Hours Worked: ")
	fmt.Scan(&hoursWorked)

	employee := NewPartTimeEmployee(id, name, hourlyRate, hoursWorked)

	if err := ui.company.AddEmployee(employee); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Part-Time Employee '%s' added successfully!\n", name)
	}
}

func (ui *CompanyUI) handleListAllEmployees() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All Employees")
	fmt.Println(strings.Repeat("=", 60))

	employees := ui.company.GetAllEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for i, emp := range employees {
		fmt.Printf("%d. %s\n", i+1, emp.GetDetails())
		fmt.Printf("   %s\n", emp.GetSalaryInfo())
		fmt.Println()
	}
}

func (ui *CompanyUI) handleListFullTimeEmployees() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("Full-Time Employees")
	fmt.Println(strings.Repeat("=", 60))

	employees := ui.company.GetFullTimeEmployees()

	if len(employees) == 0 {
		fmt.Println("No full-time employees found.")
		return
	}

	for i, emp := range employees {
		fmt.Printf("%d. %s\n", i+1, emp.GetDetails())
		fmt.Printf("   %s\n", emp.GetSalaryInfo())
		fmt.Printf("   Benefits: %s\n", emp.GetBenefits())
		fmt.Println()
	}
}

func (ui *CompanyUI) handleListPartTimeEmployees() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("Part-Time Employees")
	fmt.Println(strings.Repeat("=", 60))

	employees := ui.company.GetPartTimeEmployees()

	if len(employees) == 0 {
		fmt.Println("No part-time employees found.")
		return
	}

	for i, emp := range employees {
		fmt.Printf("%d. %s\n", i+1, emp.GetDetails())
		fmt.Printf("   %s\n", emp.GetSalaryInfo())
		fmt.Println()
	}
}

func (ui *CompanyUI) handleSearchEmployee() {
	fmt.Println("\nSearch Employee")

	var id uint64
	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	employee, err := ui.company.GetEmployee(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("\nEmployee Details")
	fmt.Println(employee.GetDetails())
	fmt.Println(employee.GetSalaryInfo())

	if ft, ok := employee.(*FullTimeEmployee); ok {
		fmt.Printf("Benefits: %s\n", ft.GetBenefits())
		fmt.Printf("Annual Salary: $%.2f\n", ft.CalculateAnnualSalary())
	} else if pt, ok := employee.(*PartTimeEmployee); ok {
		fmt.Printf("Max Hours: %.2f\n", pt.GetMaxHours())
	}
}

func (ui *CompanyUI) handleRemoveEmployee() {
	fmt.Println("\nRemove Employee")

	var id uint64
	fmt.Print("Enter Employee ID to remove: ")
	fmt.Scan(&id)

	employee, err := ui.company.GetEmployee(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if err := ui.company.RemoveEmployee(id); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Employee '%s' removed successfully!\n", employee.GetName())
	}
}

func (ui *CompanyUI) handlePayrollSummary() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("Payroll Summary")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Printf("Company: %s\n", ui.company.GetCompanyName())
	fmt.Printf("Total Employees: %d\n", ui.company.GetEmployeeCount())
	fmt.Printf("Full-Time Employees: %d\n", ui.company.GetFullTimeCount())
	fmt.Printf("Part-Time Employees: %d\n", ui.company.GetPartTimeCount())
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Total Monthly Payroll: $%.2f\n", ui.company.CalculateTotalPayroll())
	fmt.Printf("Average Employee Salary: $%.2f\n", ui.company.GetAverageSalary())

	if highest, salary := ui.company.GetHighestPaidEmployee(); highest != nil {
		fmt.Printf("Highest Paid Employee: %s ($%.2f/month)\n",
			highest.GetName(), salary)
	}
	fmt.Println()
}

func (ui *CompanyUI) readLine() string {
	ui.scanner.Scan()
	return strings.TrimSpace(ui.scanner.Text())
}

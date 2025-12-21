package company

import "fmt"

type FullTimeEmployee struct {
	ID            uint64
	Name          string
	Department    string
	MonthlySalary float64
	Benefits      string
}

func NewFullTimeEmployee(id uint64, name, department string, salary float64) *FullTimeEmployee {
	return &FullTimeEmployee{
		ID:            id,
		Name:          name,
		Department:    department,
		MonthlySalary: salary,
		Benefits:      "Health Insurance, Paid Leave, Retirement Plan",
	}
}

func (f *FullTimeEmployee) GetID() uint64 {
	return f.ID
}

func (f *FullTimeEmployee) GetName() string {
	return f.Name
}

func (f *FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("[ID: %d] %s - Full-Time Employee",
		f.ID, f.Name)
}

func (f *FullTimeEmployee) GetSalaryInfo() string {
	return fmt.Sprintf("Department: %s | Monthly Salary: $%.2f",
		f.Department, f.MonthlySalary)
}

func (f *FullTimeEmployee) CalculateMonthlySalary() float64 {
	return f.MonthlySalary
}

func (f *FullTimeEmployee) CalculateAnnualSalary() float64 {
	return f.MonthlySalary * 12
}

func (f *FullTimeEmployee) GetDepartment() string {
	return f.Department
}

func (f *FullTimeEmployee) GetBenefits() string {
	return f.Benefits
}

func (f *FullTimeEmployee) SetDepartment(department string) {
	f.Department = department
}

func (f *FullTimeEmployee) GiveRaise(percentage float64) {
	f.MonthlySalary += f.MonthlySalary * (percentage / 100)
}

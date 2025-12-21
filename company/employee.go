package company


type Employee interface {
	GetID() uint64
	GetName() string
	GetDetails() string
	GetSalaryInfo() string
	CalculateMonthlySalary() float64
}
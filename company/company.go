package company

import (
	"fmt"
	"sort"
)

type Company struct {
	name      string
	employees map[uint64]Employee
}

func NewCompany(name string) *Company {
	return &Company{
		name:      name,
		employees: make(map[uint64]Employee),
	}
}

func (c *Company) AddEmployee(employee Employee) error {
	if _, exists := c.employees[employee.GetID()]; exists {
		return fmt.Errorf("employee with ID %d already exists", employee.GetID())
	}

	c.employees[employee.GetID()] = employee
	return nil
}

func (c *Company) RemoveEmployee(id uint64) error {
	if _, exists := c.employees[id]; !exists {
		return fmt.Errorf("employee with ID %d not found", id)
	}

	delete(c.employees, id)
	return nil
}

func (c *Company) GetEmployee(id uint64) (Employee, error) {
	employee, exists := c.employees[id]
	if !exists {
		return nil, fmt.Errorf("employee with ID %d not found", id)
	}
	return employee, nil
}

func (c *Company) GetAllEmployees() []Employee {
	employees := make([]Employee, 0, len(c.employees))

	ids := make([]uint64, 0, len(c.employees))
	for id := range c.employees {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	for _, id := range ids {
		employees = append(employees, c.employees[id])
	}

	return employees
}

func (c *Company) GetEmployeeCount() int {
	return len(c.employees)
}

func (c *Company) GetCompanyName() string {
	return c.name
}

func (c *Company) CalculateTotalPayroll() float64 {
	total := 0.0
	for _, employee := range c.employees {
		total += employee.CalculateMonthlySalary()
	}
	return total
}

func (c *Company) GetFullTimeEmployees() []*FullTimeEmployee {
	fullTime := make([]*FullTimeEmployee, 0)

	for _, emp := range c.employees {
		if ft, ok := emp.(*FullTimeEmployee); ok {
			fullTime = append(fullTime, ft)
		}
	}

	return fullTime
}

func (c *Company) GetPartTimeEmployees() []*PartTimeEmployee {
	partTime := make([]*PartTimeEmployee, 0)

	for _, emp := range c.employees {
		if pt, ok := emp.(*PartTimeEmployee); ok {
			partTime = append(partTime, pt)
		}
	}

	return partTime
}

func (c *Company) GetFullTimeCount() int {
	return len(c.GetFullTimeEmployees())
}

func (c *Company) GetPartTimeCount() int {
	return len(c.GetPartTimeEmployees())
}

func (c *Company) GetAverageSalary() float64 {
	if len(c.employees) == 0 {
		return 0
	}
	return c.CalculateTotalPayroll() / float64(len(c.employees))
}

func (c *Company) GetHighestPaidEmployee() (Employee, float64) {
	if len(c.employees) == 0 {
		return nil, 0
	}

	var highest Employee
	highestSalary := 0.0

	for _, emp := range c.employees {
		salary := emp.CalculateMonthlySalary()
		if salary > highestSalary {
			highest = emp
			highestSalary = salary
		}
	}

	return highest, highestSalary
}

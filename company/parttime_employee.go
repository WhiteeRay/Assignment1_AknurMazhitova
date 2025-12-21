package company

import "fmt"

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  float64
	HoursWorked float64
	MaxHours    float64
}

func NewPartTimeEmployee(id uint64, name string, hourlyRate, hoursWorked float64) *PartTimeEmployee {
	return &PartTimeEmployee{
		ID:          id,
		Name:        name,
		HourlyRate:  hourlyRate,
		HoursWorked: hoursWorked,
		MaxHours:    40.0,
	}
}

func (p *PartTimeEmployee) GetID() uint64 {
	return p.ID
}

func (p *PartTimeEmployee) GetName() string {
	return p.Name
}

func (p *PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("[ID: %d] %s - Part-Time Employee",
		p.ID, p.Name)
}

func (p *PartTimeEmployee) GetSalaryInfo() string {
	return fmt.Sprintf("Hourly Rate: $%.2f | Hours Worked: %.2f | Monthly Salary: $%.2f",
		p.HourlyRate, p.HoursWorked, p.CalculateMonthlySalary())
}

func (p *PartTimeEmployee) CalculateMonthlySalary() float64 {
	return p.HourlyRate * p.HoursWorked
}

func (p *PartTimeEmployee) GetHourlyRate() float64 {
	return p.HourlyRate
}

func (p *PartTimeEmployee) GetHoursWorked() float64 {
	return p.HoursWorked
}

func (p *PartTimeEmployee) AddHours(hours float64) error {
	newTotal := p.HoursWorked + hours
	if newTotal > p.MaxHours {
		return fmt.Errorf("cannot exceed maximum hours of %.2f", p.MaxHours)
	}
	p.HoursWorked = newTotal
	return nil
}

func (p *PartTimeEmployee) SetHourlyRate(rate float64) {
	p.HourlyRate = rate
}

func (p *PartTimeEmployee) ResetHours() {
	p.HoursWorked = 0
}

func (p *PartTimeEmployee) GetMaxHours() float64 {
	return p.MaxHours
}

func (p *PartTimeEmployee) SetMaxHours(hours float64) {
	p.MaxHours = hours
}

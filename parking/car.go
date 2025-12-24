package ParkingLot

import "errors"

type Car struct {
	Plate       string
	Owner       string
	HoursParked int
}

type ParkingLot struct {
	Cars map[string]Car
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		Cars: make(map[string]Car),
	}
}

func (p *ParkingLot) ParkCar(car Car) error {
	if car.Plate == "" {
		return errors.New("plate must not be empty")
	}
	if car.Owner == "" {
		return errors.New("owner must not be empty")
	}
	if _, exists := p.Cars[car.Plate]; exists {
		return errors.New("car with this plate already exists")
	}

	p.Cars[car.Plate] = car
	return nil
}

func (p *ParkingLot) UnparkCar(plate string) (Car, error) {
	car, exists := p.Cars[plate]
	if !exists {
		return Car{}, errors.New("car not found")
	}

	delete(p.Cars, plate)
	return car, nil
}

func (p *ParkingLot) UpdateHours(plate string, hours int) error {
	if hours < 0 {
		return errors.New("hours must be >= 0")
	}

	car, exists := p.Cars[plate]
	if !exists {
		return errors.New("car not found")
	}

	car.HoursParked = hours
	p.Cars[plate] = car
	return nil
}

func (p *ParkingLot) ListCars() []Car {
	cars := []Car{}
	for _, car := range p.Cars {
		cars = append(cars, car)
	}
	return cars
}

func (p *ParkingLot) Bills(ratePerHour int) map[string]int {
	bills := make(map[string]int)

	for plate, car := range p.Cars {
		bills[plate] = car.HoursParked * ratePerHour
	}

	return bills
}

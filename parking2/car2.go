package parking2


import "errors"

type Car struct{
	Plate string
	Owner string
	HoursParked int
}

type ParkingLot struct{
	Cars map[string]Car;
}

func NewParkingLot() *ParkingLot{
	return &ParkingLot{
		Cars:make(map[string]Car),
	}
}
func(p* ParkingLot)ParkCar(car Car) error{
	if car.Plate == ""{
		return errors.New("Plate must not be empty")

	}
	if car.Owner ==""{
		return errors.New("Owner should be unique")
	}

	if _,exists:= p.Cars[car.Plate]; exists{
		return errors.New("Plate must be unique")
	}
	p.Cars[car.Plate] = car
	return nil
}

func (p *ParkingLot)UnparkCar(plate string)(Car,error){
	car,exists := p.Cars[plate]
	if !exists{
		return Car{}, errors.New("Car must exist")
	}
	delete(p.Cars, plate)
	return car,nil
}

func (p *ParkingLot) UpdateHours(plate string, hours int) error{
	car,exists :=p.Cars[plate]
	if !exists{
		return errors.New("Car must exist")
	}
	if hours < 0{
		return errors.New("hours must be>=0")
	}
	car.HoursParked = hours
	p.Cars[plate] = car
	return nil
}

func (p ParkingLot) ListCars() []Car{
	cars := []Car{}
	for _,v:= range p.Cars{
		cars = append(cars, v)
	}
	return cars
}

func(p *ParkingLot) Bills(ratePerHour int) map[string]int{
	bill:=make(map[string]int)
	for k,v:=range p.Cars{
		cost:=v.HoursParked * ratePerHour
		bill[k] = cost
	}
	return bill
}


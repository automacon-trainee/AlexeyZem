package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlat string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		space: list.New(),
	}
}

func (p *ParkingLot) Park(c Car) {
	p.space.PushBack(c)
}

func (p *ParkingLot) Leave() {
	if p.space.Len() == 0 {
		fmt.Println("Парковка пуста")
		return
	}
	car := p.space.Front()
	fmt.Println("Автомобиль", car.Value.(Car).LicensePlat, "покинул парковку")
	p.space.Remove(car)
}

func main() {
	parkingLot := NewParkingLot()
	parkingLot.Park(Car{LicensePlat: "ABC-123"})
	parkingLot.Park(Car{LicensePlat: "XYZ-789"})
	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
}

package main

import (
	"fmt"
	"lld-examples/internal/parkinglot/entity"
	"lld-examples/internal/parkinglot/service"
	"time"
)

func main() {

	// Create spots
	spotsFloor1 := []*entity.ParkingSpot{
		{ID: 1, Type: entity.Car},
		{ID: 2, Type: entity.Motorbike},
		{ID: 3, Type: entity.Truck},
	}

	floor1 := &entity.ParkingFloor{
		FloorNumber: 1,
		Spots:       spotsFloor1,
	}

	parkingLot := service.NewParkingLot([]*entity.ParkingFloor{floor1})

	vehicle := &entity.Vehicle{
		LicensePlate: "KA-01-1234",
		Type:         entity.Car,
	}

	ticket, err := parkingLot.ParkVehicle(vehicle)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Vehicle parked. Ticket ID:", ticket.ID)

	time.Sleep(2 * time.Second)

	amount, err := parkingLot.ExitVehicle(ticket.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parking fee:", amount)
}

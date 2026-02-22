package entity

import "errors"

type ParkingSpot struct {
	ID       int
	Type     VehicleType
	Occupied bool
	Vehicle  *Vehicle
}

func (ps *ParkingSpot) Park(vehicle *Vehicle) error {
	if ps.Occupied {
		return errors.New("spot already occupied")
	}
	if ps.Type != vehicle.Type {
		return errors.New("vehicle type mismatch")
	}

	ps.Vehicle = vehicle
	ps.Occupied = true
	return nil
}

func (ps *ParkingSpot) Leave() {
	ps.Vehicle = nil
	ps.Occupied = false
}

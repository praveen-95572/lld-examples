package entity

import (
	"fmt"
	"time"
)

type Ticket struct {
	ID        string
	Vehicle   *Vehicle
	Spot      *ParkingSpot
	EntryTime time.Time
}

func NewTicket(vehicle *Vehicle, spot *ParkingSpot) *Ticket {
	return &Ticket{
		ID:        fmt.Sprintf("%s-%d", vehicle.LicensePlate, time.Now().UnixNano()),
		Vehicle:   vehicle,
		Spot:      spot,
		EntryTime: time.Now(),
	}
}

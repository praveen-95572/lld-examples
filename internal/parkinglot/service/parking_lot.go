package service

import (
	"errors"
	"lld-examples/internal/parkinglot/entity"
	"log"
	"sync"
	"time"
)

var (
	instance *ParkingLot
	once     sync.Once
)

type ParkingLot struct {
	Floors  []*entity.ParkingFloor
	Tickets map[string]*entity.Ticket
}

func initializeParkingLot(floors []*entity.ParkingFloor) {
	instance = &ParkingLot{
		Floors:  floors,
		Tickets: make(map[string]*entity.Ticket),
	}
}

func NewParkingLot(floors []*entity.ParkingFloor) *ParkingLot {
	once.Do(func() {
		log.Println("Will create only one instance")
		initializeParkingLot(floors)
	})
	return instance
}

func (pl *ParkingLot) ParkVehicle(vehicle *entity.Vehicle) (*entity.Ticket, error) {

	for _, floor := range pl.Floors {
		spot := floor.FindAvailableSpot(vehicle.Type)
		if spot != nil {
			err := spot.Park(vehicle)
			if err != nil {
				return nil, err
			}

			ticket := entity.NewTicket(vehicle, spot)
			pl.Tickets[ticket.ID] = ticket
			return ticket, nil
		}
	}

	return nil, errors.New("parking lot full")
}

func (pl *ParkingLot) ExitVehicle(ticketID string) (float64, error) {

	ticket, exists := pl.Tickets[ticketID]
	if !exists {
		return 0, errors.New("invalid ticket")
	}

	duration := time.Since(ticket.EntryTime).Hours()
	hours := int(duration) + 1

	rate := 10.0 // flat rate per hour
	amount := float64(hours) * rate

	ticket.Spot.Leave()
	delete(pl.Tickets, ticketID)

	return amount, nil
}

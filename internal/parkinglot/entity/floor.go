package entity

type ParkingFloor struct {
	FloorNumber int
	Spots       []*ParkingSpot
}

func (pf *ParkingFloor) FindAvailableSpot(vehicleType VehicleType) *ParkingSpot {
	for _, spot := range pf.Spots {
		if !spot.Occupied && spot.Type == vehicleType {
			return spot
		}
	}
	return nil
}

package entity

type VehicleType string

const (
	Car       VehicleType = "Car"
	Motorbike VehicleType = "Motorbike"
	Truck     VehicleType = "Truck"
)

type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

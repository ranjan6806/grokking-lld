package models

type VehicleType string

const (
	Car        VehicleType = "car"
	Truck      VehicleType = "truck"
	Van        VehicleType = "van"
	Motorcycle VehicleType = "motorcycle"
)

type VehicleInterface interface {
	GetVehicleType() VehicleType
	GetLicensePlate() string
}

type Vehicle struct {
	VehicleType  VehicleType
	LicensePlate string
}

func (v *Vehicle) GetVehicleType() VehicleType {
	return v.VehicleType
}

func (v *Vehicle) GetLicensePlate() string {
	return v.LicensePlate
}

// Factory function returns VehicleInterface
func NewVehicle(vehicleType VehicleType, licensePlate string) VehicleInterface {
	return &Vehicle{VehicleType: vehicleType, LicensePlate: licensePlate}
}

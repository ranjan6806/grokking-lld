package controller

import "parking-lot/models"

type VehicleControllerInterface interface {
	CreateVehicle(vehicleType models.VehicleType, licensePlate string) models.VehicleInterface
}

type VehicleController struct{}

func NewVehicleController() *VehicleController {
	return &VehicleController{}
}

func (vc *VehicleController) CreateVehicle(vehicleType models.VehicleType, licensePlate string) models.VehicleInterface {
	return &models.Vehicle{
		VehicleType:  vehicleType,
		LicensePlate: licensePlate,
	}
}

package controller

import (
	"parking-lot/models"
	"parking-lot/service"
)

type ParkingControllerInterface interface {
	IssueTicket(vehicle models.VehicleInterface, payment models.PaymentInterface) (models.TicketInterface, error)
	ProcessExit(ticket models.TicketInterface) error
	ShowFreeSpots() map[models.SpotType]int
	GetCapacity(spotType models.SpotType) int
	GetCurrentUsage(spotType models.SpotType) int
}

type ParkingController struct {
	parkingService service.ParkingLotServiceInterface
	paymentService service.PaymentServiceInterface
}

func (pc *ParkingController) GetCapacity(spotType models.SpotType) int {
	return pc.parkingService.GetCapacity(spotType)
}

func (pc *ParkingController) GetCurrentUsage(spotType models.SpotType) int {
	return pc.parkingService.GetCurrentUsage(spotType)
}

func (pc *ParkingController) IssueTicket(vehicle models.VehicleInterface, payment models.PaymentInterface) (models.TicketInterface, error) {
	return pc.parkingService.IssueTicket(vehicle, payment)
}

func (pc *ParkingController) ProcessExit(
	ticket models.TicketInterface,
) error {
	return pc.parkingService.ProcessExit(ticket)
}

func (pc *ParkingController) ShowFreeSpots() map[models.SpotType]int {
	return pc.parkingService.ShowFreeSpots()
}

func NewParkingController(service service.ParkingLotServiceInterface) ParkingControllerInterface {
	return &ParkingController{
		parkingService: service,
	}
}

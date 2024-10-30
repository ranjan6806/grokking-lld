package controller

import (
	"parking-lot/models"
	"parking-lot/service"
)

type ParkingControllerInterface interface {
	IssueTicket(vehicle models.VehicleInterface) (models.TicketInterface, error)
	ProcessExit(ticket models.TicketInterface, method models.PaymentMethod, hourlyRate float64) (models.PaymentInterface, error)
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

func (pc *ParkingController) IssueTicket(vehicle models.VehicleInterface) (models.TicketInterface, error) {
	return pc.parkingService.IssueTicket(vehicle)
}

func (pc *ParkingController) ProcessExit(
	ticket models.TicketInterface,
	method models.PaymentMethod,
	hourlyRate float64,
) (models.PaymentInterface, error) {
	return pc.parkingService.ProcessExit(ticket, method, hourlyRate)
}

func (pc *ParkingController) ShowFreeSpots() map[models.SpotType]int {
	return pc.parkingService.ShowFreeSpots()
}

func NewParkingController(service service.ParkingLotServiceInterface) ParkingControllerInterface {
	return &ParkingController{
		parkingService: service,
	}
}

package controller

import (
	"parking-lot/models"
	"parking-lot/service"
)

type ParkingControllerInterface interface {
	IssueTicket(vehicle models.VehicleInterface) (models.TicketInterface, error)
	ProcessExit(ticket models.TicketInterface, method models.PaymentMethod, hourlyRate float64) (models.PaymentInterface, error)
	ShowFreeSpots() map[models.SpotType]int
}

type ParkingController struct {
	parkingService service.ParkingLotServiceInterface
	paymentService service.PaymentServiceInterface
}

func NewParkingController(service service.ParkingLotServiceInterface) ParkingControllerInterface {
	return &ParkingController{
		parkingService: service,
	}
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

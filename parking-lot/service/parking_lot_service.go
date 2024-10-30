package service

import (
	"errors"
	"parking-lot/models"
	"time"
)

type ParkingLotServiceInterface interface {
	IssueTicket(vehicle models.VehicleInterface) (models.TicketInterface, error)
	ProcessExit(ticket models.TicketInterface, method models.PaymentMethod, hourlyRate float64) (models.PaymentInterface, error)
	ShowFreeSpots() map[models.SpotType]int
	GetCapacity(spotType models.SpotType) int
	GetCurrentUsage(spotType models.SpotType) int
}

type ParkingLotService struct {
	spots        map[models.SpotType][]models.SpotInterface
	displayBoard DisplayBoardServiceInterface
	capacities   map[models.SpotType]int
	currentUsage map[models.SpotType]int
}

func NewParkingLotService(
	spots map[models.SpotType][]models.SpotInterface,
	displayBoard DisplayBoardServiceInterface,
	capacities map[models.SpotType]int,
) ParkingLotServiceInterface {
	// Initialize currentUsage based on initial free spots
	currentUsage := make(map[models.SpotType]int)
	for spotType, spotList := range spots {
		currentUsage[spotType] = 0
		displayBoard.SetFreeSpots(spotType, len(spotList))
	}

	pl := &ParkingLotService{
		spots:        spots,
		displayBoard: displayBoard,
		capacities:   capacities,
		currentUsage: currentUsage,
	}

	return pl
}

func (pl *ParkingLotService) GetCapacity(spotType models.SpotType) int {
	return pl.capacities[spotType]
}

func (pl *ParkingLotService) GetCurrentUsage(spotType models.SpotType) int {
	return pl.currentUsage[spotType]
}

func (pl *ParkingLotService) IssueTicket(vehicle models.VehicleInterface) (models.TicketInterface, error) {
	spotType := getSpotTypeForVehicle(vehicle.GetVehicleType())

	if pl.currentUsage[spotType] >= pl.capacities[spotType] {
		return nil, errors.New("no available spot for vehicle type")
	}

	// Find a free spot of the required spot type
	for _, spot := range pl.spots[spotType] {
		if spot.IsFree() {
			// Park the vehicle
			spot.ParkVehicle(vehicle)
			pl.currentUsage[spotType]++
			pl.displayBoard.DecrementFreeSpots(spotType)
			return models.NewTicket(vehicle, spot), nil
		}
	}

	return nil, errors.New("no available spot for vehicle")
}

func (pl *ParkingLotService) ProcessExit(ticket models.TicketInterface, method models.PaymentMethod, hourlyRate float64) (models.PaymentInterface, error) {
	duration := calculatePayment(ticket, hourlyRate)

	ticket.GetSpot().RemoveVehicle()
	pl.currentUsage[ticket.GetSpot().GetSpotType()]--
	pl.displayBoard.IncrementFreeSpots(ticket.GetSpot().GetSpotType())

	payment := &models.Payment{
		Ticket:        ticket,
		Amount:        duration,
		PaymentMethod: method,
	}

	return payment, nil
}

func (pl *ParkingLotService) ShowFreeSpots() map[models.SpotType]int {
	return pl.displayBoard.ShowFreeSpots()
}

func getSpotTypeForVehicle(vehicleType models.VehicleType) models.SpotType {
	switch vehicleType {
	case models.Car, models.Van:
		return models.Compact
	case models.Motorcycle:
		return models.MotorcycleSpot
	case models.Truck:
		return models.Large
	default:
		return models.Compact
	}
}

func calculatePayment(ticket models.TicketInterface, hourlyRate float64) float64 {
	duration := time.Since(ticket.GetEntryTime()).Hours()
	if duration < 1 {
		return hourlyRate
	}

	return float64(int(duration+0.99)) * hourlyRate
}

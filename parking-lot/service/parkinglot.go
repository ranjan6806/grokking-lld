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
}

type ParkingLotService struct {
	capacity        int
	currentVehicles int
	spots           map[models.SpotType][]models.SpotInterface
	displayBoard    DisplayBoardServiceInterface
}

func NewParkingLotService(spots map[models.SpotType][]models.SpotInterface, displayBoard DisplayBoardServiceInterface) ParkingLotServiceInterface {
	pl := &ParkingLotService{
		spots:        spots,
		displayBoard: displayBoard,
	}

	pl.initializeSpots()
	return pl
}

func (pl *ParkingLotService) initializeSpots() {
	spotCounts := map[models.SpotType]int{
		models.Handicapped:    1000,
		models.Compact:        15000,
		models.Large:          12000,
		models.MotorcycleSpot: 12000,
	}

	for spotType, count := range spotCounts {
		for i := 0; i < count; i++ {
			spot := &models.Spot{
				ID:       i,
				SpotType: spotType,
				Free:     true,
				Vehicle:  nil,
			}

			pl.spots[spotType] = append(pl.spots[spotType], spot)
		}
		pl.displayBoard.SetFreeSpots(spotType, count)
	}
}

func (pl *ParkingLotService) IssueTicket(vehicle models.VehicleInterface) (models.TicketInterface, error) {
	if pl.currentVehicles >= pl.capacity {
		return nil, errors.New("parking lot is full")
	}

	spotType := GetSpotTypeForVehicle(vehicle.GetVehicleType())
	assignedSpot, err := pl.findFreeSpot(spotType)
	if err != nil {
		return nil, err
	}

	assignedSpot.ParkVehicle(vehicle)
	pl.currentVehicles++
	pl.displayBoard.DecrementFreeSpots(spotType)

	return models.NewTicket(vehicle, assignedSpot), nil
}

func (pl *ParkingLotService) ProcessExit(ticket models.TicketInterface, method models.PaymentMethod, hourlyRate float64) (models.PaymentInterface, error) {
	duration := CalculatePayment(ticket, hourlyRate)

	ticket.GetSpot().RemoveVehicle()
	pl.currentVehicles--
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

func (pl *ParkingLotService) findFreeSpot(spotType models.SpotType) (models.SpotInterface, error) {
	spots, exists := pl.spots[spotType]
	if !exists {
		return nil, errors.New("spot not exists")
	}

	for _, spot := range spots {
		if spot.IsFree() {
			return spot, nil
		}
	}

	return nil, errors.New("no available spot for vehicle type")
}

func GetSpotTypeForVehicle(vehicleType models.VehicleType) models.SpotType {
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

func CalculatePayment(ticket models.TicketInterface, hourlyRate float64) float64 {
	duration := time.Since(ticket.GetEntryTime()).Hours()
	if duration < 1 {
		return hourlyRate
	}

	return float64(int(duration+0.99)) * hourlyRate
}

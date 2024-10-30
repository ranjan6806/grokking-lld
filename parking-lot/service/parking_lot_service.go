package service

import (
	"errors"
	"parking-lot/models"
)

type ParkingLotServiceInterface interface {
	IssueTicket(vehicle models.VehicleInterface, payment models.PaymentInterface) (models.TicketInterface, error)
	ProcessExit(ticket models.TicketInterface) error
	ShowFreeSpots() map[models.SpotType]int
	GetCapacity(spotType models.SpotType) int
	GetCurrentUsage(spotType models.SpotType) int
}

type ParkingLotService struct {
	spotRepo       models.SpotRepositoryInterface
	ticketRepo     models.TicketRepositoryInterface
	displayBoard   DisplayBoardServiceInterface
	paymentService PaymentServiceInterface
	capacities     map[models.SpotType]int
	currentUsage   map[models.SpotType]int
}

func (pl *ParkingLotService) GetCapacity(spotType models.SpotType) int {
	return pl.capacities[spotType]
}

func (pl *ParkingLotService) GetCurrentUsage(spotType models.SpotType) int {
	return pl.currentUsage[spotType]
}

func (pl *ParkingLotService) IssueTicket(vehicle models.VehicleInterface, payment models.PaymentInterface) (models.TicketInterface, error) {
	spotType := getSpotTypeForVehicle(vehicle.GetVehicleType())

	if pl.currentUsage[spotType] >= pl.capacities[spotType] {
		return nil, errors.New("no available spot for vehicle type")
	}

	// Process payment (assuming each parking spot costs a fixed amount, e.g., $10)
	paymentAmount := 10.0 // This could be dynamic based on spot type or duration
	success, err := pl.paymentService.ProcessPayment(payment, paymentAmount)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errors.New("payment processing failed")
	}

	// Find a free spot of the required spot type
	spots, _ := pl.spotRepo.GetAllSpots()

	for _, spot := range spots {
		if spot.GetSpotType() == spotType && spot.IsFree() {
			// Park the vehicle
			spot.ParkVehicle(vehicle)
			pl.currentUsage[spotType]++
			pl.displayBoard.DecrementFreeSpots(spotType)

			ticket := models.NewTicket(vehicle, spot, payment)
			err = pl.ticketRepo.CreateTicket(ticket)
			if err != nil {
				// Rollback spot occupation
				spot.RemoveVehicle()
				pl.spotRepo.UpdateSpot(spot)
				pl.currentUsage[spotType]--
				pl.displayBoard.IncrementFreeSpots(spotType)
				return nil, err
			}

			return ticket, nil
		}
	}

	return nil, errors.New("no available spot for vehicle")
}

func (pl *ParkingLotService) ProcessExit(ticket models.TicketInterface) error {
	spot := ticket.GetSpot()
	spotType := spot.GetSpotType()

	if spot.IsFree() {
		return errors.New("spot is free")
	}

	spot.RemoveVehicle()
	err := pl.spotRepo.UpdateSpot(spot)
	if err != nil {
		return err
	}

	pl.currentUsage[ticket.GetSpot().GetSpotType()]--
	pl.displayBoard.IncrementFreeSpots(spotType)

	err = pl.ticketRepo.DeleteTicket(ticket.GetID())
	if err != nil {
		return err
	}

	return nil
}

func (pl *ParkingLotService) ShowFreeSpots() map[models.SpotType]int {
	return pl.displayBoard.ShowFreeSpots()
}

func NewParkingLotService(
	spotRepo models.SpotRepositoryInterface,
	ticketRepo models.TicketRepositoryInterface,
	spots map[models.SpotType][]models.SpotInterface,
	displayBoard DisplayBoardServiceInterface,
	capacities map[models.SpotType]int,
	paymentService PaymentServiceInterface,
) ParkingLotServiceInterface {
	// Initialize currentUsage based on initial free spots
	currentUsage := make(map[models.SpotType]int)
	for spotType, spotList := range spots {
		currentUsage[spotType] = 0
		displayBoard.SetFreeSpots(spotType, len(spotList))
	}

	pl := &ParkingLotService{
		spotRepo:       spotRepo,
		ticketRepo:     ticketRepo,
		displayBoard:   displayBoard,
		paymentService: paymentService,
		capacities:     capacities,
		currentUsage:   currentUsage,
	}

	return pl
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

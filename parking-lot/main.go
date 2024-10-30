package main

import (
	"fmt"
	"parking-lot/controller"
	"parking-lot/models"
	"parking-lot/service"
)

func main() {
	// Create a map of spots for simplicity
	spots := map[models.SpotType][]models.SpotInterface{
		models.Compact:        {models.NewSpot(1, models.Compact), models.NewSpot(2, models.Compact)},
		models.Large:          {models.NewSpot(3, models.Large)},
		models.MotorcycleSpot: {models.NewSpot(4, models.MotorcycleSpot)},
	}

	// Create services and controllers using interfaces
	displayBoard := service.NewDisplayBoardService()
	parkingLotService := service.NewParkingLotService(spots, displayBoard)
	parkingController := controller.NewParkingController(parkingLotService)

	// Use controller methods
	vehicle := models.NewVehicle(models.Car, "ABC123")
	ticket, err := parkingController.IssueTicket(vehicle)
	if err != nil {
		fmt.Println("Error issuing ticket:", err)
		return
	}

	fmt.Printf("Ticket issued: %+v\n", ticket)

	freeSpots := parkingController.ShowFreeSpots()
	fmt.Println("Free spots after issuing ticket:", freeSpots)

	// Process vehicle exit
	p, err := parkingController.ProcessExit(ticket, models.CreditCard, 2.5)
	if err != nil {
		fmt.Println("Error processing exit:", err)
		return
	}

	fmt.Println("payment p: -", p)

	freeSpots = parkingController.ShowFreeSpots()
	fmt.Println("Free spots after exit:", freeSpots)
}

package main

import (
	"fmt"
	"parking-lot/controller"
	"parking-lot/models"
	"parking-lot/service"
)

func main() {
	// Define spots for each SpotType
	spots := map[models.SpotType][]models.SpotInterface{
		models.Compact:        {models.NewSpot(1, models.Compact), models.NewSpot(2, models.Compact), models.NewSpot(3, models.Compact)},
		models.Large:          {models.NewSpot(4, models.Large), models.NewSpot(5, models.Large)},
		models.MotorcycleSpot: {models.NewSpot(6, models.MotorcycleSpot)},
		models.Handicapped:    {models.NewSpot(7, models.Handicapped)},
	}

	// Define capacities for each SpotType
	capacities := map[models.SpotType]int{
		models.Compact:        3, // Number of compact spots
		models.Large:          2, // Number of large spots
		models.MotorcycleSpot: 1, // Number of motorcycle spots
		models.Handicapped:    1, // Number of handicapped spots
	}

	// Initialize DisplayBoardService
	displayBoard := service.NewDisplayBoardService()

	// Initialize PaymentService
	paymentService := service.NewPaymentService()

	// Initialize ParkingLotService with per-SpotType capacities and PaymentService
	parkingLotService := service.NewParkingLotService(spots, displayBoard, capacities, paymentService)

	// Initialize Controllers
	parkingController := controller.NewParkingController(parkingLotService)
	displayController := controller.NewDisplayController(displayBoard)
	vehicleController := controller.NewVehicleController()

	// Create vehicles using VehicleController
	vehicle1 := vehicleController.CreateVehicle(models.Car, "ABC123")
	vehicle2 := vehicleController.CreateVehicle(models.Truck, "XYZ789")
	vehicle3 := vehicleController.CreateVehicle(models.Van, "VAN456")
	vehicle4 := vehicleController.CreateVehicle(models.Motorcycle, "MOTO321")
	vehicle5 := vehicleController.CreateVehicle(models.Car, "CAR999") // This should exceed capacity for Compact spots

	// Create payments
	payment1, err := models.NewPayment(models.Cash)
	if err != nil {
		fmt.Printf("Error creating payment1: %v\n", err)
		return
	}

	payment2, err := models.NewPayment(models.CreditCard, "123456789012", "12/25", "123")
	if err != nil {
		fmt.Printf("Error creating payment2: %v\n", err)
		return
	}

	payment3, err := models.NewPayment(models.CreditCard, "987654321098", "11/24", "456")
	if err != nil {
		fmt.Printf("Error creating payment3: %v\n", err)
		return
	}

	payment4, err := models.NewPayment(models.Cash)
	if err != nil {
		fmt.Printf("Error creating payment4: %v\n", err)
		return
	}

	payment5, err := models.NewPayment(models.CreditCard, "111122223333", "10/23", "789")
	if err != nil {
		fmt.Printf("Error creating payment5: %v\n", err)
		return
	}

	// Issue tickets with payments
	tickets := []models.TicketInterface{}
	vehicles := []models.VehicleInterface{vehicle1, vehicle2, vehicle3, vehicle4, vehicle5}
	payments := []models.PaymentInterface{payment1, payment2, payment3, payment4, payment5}

	for i, v := range vehicles {
		payment := payments[i]
		ticket, err := parkingController.IssueTicket(v, payment)
		if err != nil {
			fmt.Printf("Error issuing ticket for vehicle %s: %v\n", v.GetLicensePlate(), err)
			continue
		}
		fmt.Printf("Ticket issued: ID=%s, Vehicle=%s, Spot=%d, EntryTime=%s, PaymentType=%s\n",
			ticket.GetID(),
			ticket.GetVehicle().GetLicensePlate(),
			ticket.GetSpot().GetID(),
			ticket.GetEntryTime().Format("2006-01-02 15:04:05"),
			ticket.GetPayment().GetPaymentType(),
		)
		tickets = append(tickets, ticket)
	}

	// Show available spots
	fmt.Println("\nFree spots after issuing tickets:")
	freeSpots := displayController.ShowFreeSpots()
	for spotType, count := range freeSpots {
		fmt.Printf("%s: %d\n", spotType, count)
	}

	// Show parking lot capacities and current usage
	fmt.Println("\nParking Lot Capacities and Current Usage:")
	for spotType, capacity := range capacities {
		usage := parkingController.GetCurrentUsage(spotType)
		fmt.Printf("%s - Capacity: %d, Current Usage: %d\n", spotType, capacity, usage)
	}

	// Process exits for some vehicles
	fmt.Println("\nProcessing exits for some vehicles.")
	if len(tickets) >= 3 {
		for i := 0; i < 3; i++ {
			err := parkingController.ProcessExit(tickets[i])
			if err != nil {
				fmt.Printf("Error processing exit for ticket %s: %v\n", tickets[i].GetID(), err)
				continue
			}
			fmt.Printf("Processed exit for ticket: %s\n", tickets[i].GetID())
		}
	}

	// Show available spots after exits
	fmt.Println("\nFree spots after processing exits:")
	freeSpots = displayController.ShowFreeSpots()
	for spotType, count := range freeSpots {
		fmt.Printf("%s: %d\n", spotType, count)
	}

	// Display updated capacities and current usage
	fmt.Println("\nUpdated Parking Lot Capacities and Current Usage:")
	for spotType, capacity := range capacities {
		usage := parkingController.GetCurrentUsage(spotType)
		fmt.Printf("%s - Capacity: %d, Current Usage: %d\n", spotType, capacity, usage)
	}
}

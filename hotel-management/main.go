package main

import (
	"fmt"
	"hotel-management/address"
	"hotel-management/booking"
	"hotel-management/hotel"
	"hotel-management/room"
	"hotel-management/service"
)

func main() {
	// Initialise hotel repo
	hotelRepo := hotel.NewHotelRepository()

	// Initialise admin service
	adminService := service.NewAdminService(hotelRepo)

	// Add hotel
	err := adminService.AddHotel("h1", "Hotel 1")
	if err != nil {
		fmt.Println("error creating hotel", err)
	}

	err = adminService.CreateHotelBranch("h1", "b1", "Branch 1", address.Address{City: "Delhi"})
	if err != nil {
		fmt.Println("error creating hotel branch", err)
	}

	err = adminService.CreateHotelBranch("h1", "b2", "Branch 2", address.Address{City: "Mumbai"})
	if err != nil {
		fmt.Println("error creating hotel branch", err)
	}

	err = adminService.AddRoom("r1", "h1", "b1", 200, room.RoomTypeStandard)
	err = adminService.AddRoom("r2", "h1", "b1", 300, room.RoomTypeStandard)

	err = adminService.AddRoom("r3", "h1", "b2", 400, room.RoomTypeStandard)
	err = adminService.AddRoom("r4", "h1", "b2", 500, room.RoomTypeStandard)

	err = adminService.ShowAllHotelDetails()
	if err != nil {
		fmt.Println("error showing hotel details", err)
	}

	// Initialise booking repository
	bookingRepo := booking.NewBookingRepository()

	// Initialise booking service
	bookingService := service.NewBookingService(bookingRepo, adminService)

	branch1, err := adminService.GetHotelBranch("h1", "b1")
	if err != nil {
		fmt.Println("error getting hotel branch", err)
	}

	err = bookingService.EnableHotelBranchBookings("h1", branch1)
	if err != nil {
		fmt.Println("error enabling hotel booking", err)
	}

	err = bookingService.BookRoom("book1", "h1", "b1", "guest1", room.RoomTypeStandard, 4, 500)
	if err != nil {
		fmt.Println("error booking room", err)
	}

	err = adminService.ShowAllHotelDetails()
	if err != nil {
		fmt.Println("error showing hotel details", err)
	}
}

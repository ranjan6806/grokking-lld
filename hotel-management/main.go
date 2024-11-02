package main

import (
	"fmt"
	"hotel-management/hotel"
	"hotel-management/hotel_branch"
	"hotel-management/room"
)

func main() {
	// Initialise repositories
	hotelRepo := hotel.NewHotelRepository()
	branchRepo := hotel_branch.NewHotelBranchRepository()

	// Create a hotel;
	hotel := hotel.NewHotel("hotel1", "Raja Hotel")
	hotelRepo.AddHotel(hotel)

	// Create a branch and add it to hotel
	branch1 := hotel_branch.NewBranch("branch1", "Main branch")
	branchRepo.AddBranch(branch1)
	hotel.AddBranch(branch1)

	branch2 := hotel_branch.NewBranch("branch2", "Second branch")
	branchRepo.AddBranch(branch2)
	hotel.AddBranch(branch2)

	// Create rooms and add them to branch
	room1 := room.NewRoom(room.RoomTypeDeluxe, "room1", room.RoomStatusAvailable, 300)
	room2 := room.NewRoom(room.RoomTypeDeluxe, "room2", room.RoomStatusAvailable, 300)
	branch1.AddRoom(room1)
	branch2.AddRoom(room2)

	// Print hotel details
	retrievedHotel, err := hotelRepo.GetHotel("hotel1")
	if err == nil {
		fmt.Printf("Hotel: %+v\n", retrievedHotel)
	}

	for _, brn := range retrievedHotel.Branches {
		fmt.Printf("brn - %+v\n", brn)

		for _, rm := range brn.Rooms {
			fmt.Printf("room - %+v\n", rm)
		}
	}

}

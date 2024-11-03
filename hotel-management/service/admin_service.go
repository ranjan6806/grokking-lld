package service

import (
	"fmt"
	"hotel-management/address"
	"hotel-management/hotel"
	"hotel-management/hotel_branch"
	"hotel-management/room"
)

type AdminService interface {
	AddHotel(hotelID, name string) error
	GetHotel(hotelID string) (*hotel.Hotel, error)
	CreateHotelBranch(hotelID, branchID, name string, address address.Address) error
	GetHotelBranch(hotelID, branchID string) (*hotel_branch.HotelBranch, error)
	GetRoom(hotelID, branchID, roomID string) (room.Room, error)
	AddRoom(roomNumber, hotelID, branchID string, bookingPrice float64, roomType room.RoomType) error
	ShowAllHotelDetails() error
}

type AdminServiceImpl struct {
	hotelRepo hotel.HotelRepository
}

func NewAdminService(hotelRepo hotel.HotelRepository) AdminService {
	return &AdminServiceImpl{
		hotelRepo: hotelRepo,
	}
}

func (s *AdminServiceImpl) AddHotel(hotelID, name string) error {
	newHotel := hotel.NewHotel(hotelID, name)
	return s.hotelRepo.AddHotel(newHotel)
}

func (s *AdminServiceImpl) GetHotel(hotelID string) (*hotel.Hotel, error) {
	return s.hotelRepo.GetHotel(hotelID)
}

func (s *AdminServiceImpl) GetHotelBranch(hotelID, branchID string) (*hotel_branch.HotelBranch, error) {
	hotelObj, err := s.hotelRepo.GetHotel(hotelID)
	if err != nil {
		return nil, err
	}

	branchObj, err := hotelObj.GetBranch(branchID)
	if err != nil {
		return nil, err
	}

	return branchObj, nil
}

func (s *AdminServiceImpl) GetRoom(hotelID, branchID, roomID string) (room.Room, error) {
	hotelObj, err := s.hotelRepo.GetHotel(hotelID)
	if err != nil {
		return nil, err
	}

	branchObj, err := hotelObj.GetBranch(branchID)
	if err != nil {
		return nil, err
	}

	roomObj, err := branchObj.GetRoom(roomID)
	if err != nil {
		return nil, err
	}

	return roomObj, nil
}

func (s *AdminServiceImpl) CreateHotelBranch(hotelID, branchID, name string, address address.Address) error {
	hotelObj, err := s.hotelRepo.GetHotel(hotelID)
	if err != nil {
		return err
	}

	branchObj := hotel_branch.NewBranch(branchID, name, address)
	err = hotelObj.AddBranch(branchObj)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminServiceImpl) AddRoom(roomNumber, hotelID, branchID string, bookingPrice float64, roomType room.RoomType) error {
	hotelObj, err := s.hotelRepo.GetHotel(hotelID)
	if err != nil {
		return err
	}

	branchObj, err := hotelObj.GetBranch(branchID)
	if err != nil {
		return err
	}

	roomObj := room.NewRoom(roomType, roomNumber, bookingPrice)
	err = branchObj.AddRoom(roomObj)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminServiceImpl) ShowAllHotelDetails() error {
	hotelObjs := s.hotelRepo.GetAllHotels()
	for _, hotelObj := range hotelObjs {
		fmt.Println("HOTEL DATA")
		fmt.Printf("Hotel ID - %s, Hotel Name - %s\n", hotelObj.HotelID, hotelObj.Name)

		for _, branchObj := range hotelObj.Branches {
			fmt.Println()
			fmt.Println("BRANCH DATA")
			fmt.Printf("Branch ID - %s, Branch Name - %s, Address - %+v\n", branchObj.BranchID, branchObj.Name, branchObj.Address)

			for _, roomObj := range branchObj.Rooms {
				fmt.Printf("Room Number - %s, Room Type - %v, Is available - %v\n", roomObj.GetRoomNumber(), roomObj.GetRoomType(), roomObj.IsAvailable())
			}
		}
	}

	return nil
}

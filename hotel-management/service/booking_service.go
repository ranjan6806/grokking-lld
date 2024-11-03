package service

import (
	"hotel-management/booking"
	"hotel-management/hotel_branch"
	"hotel-management/room"
)

type BookingService interface {
	EnableHotelBranchBookings(hotelID string, branch *hotel_branch.HotelBranch) error
	BookRoom(bookingID, hotelID, branchID, guestID string, roomType room.RoomType, duration uint32, advance float64) error
	CancelBooking(bookingID string) error
}

type BookingServiceImpl struct {
	bookingRepo  booking.BookingRepository
	adminService AdminService
}

func NewBookingService(bookingRepo booking.BookingRepository, adminService AdminService) BookingService {
	return &BookingServiceImpl{
		bookingRepo:  bookingRepo,
		adminService: adminService,
	}
}

func (s *BookingServiceImpl) EnableHotelBranchBookings(hotelID string, branch *hotel_branch.HotelBranch) error {
	return s.bookingRepo.AddBranch(hotelID, branch)
}

func (s *BookingServiceImpl) BookRoom(
	bookingID, hotelID, branchID, guestID string,
	roomType room.RoomType,
	duration uint32,
	advance float64,
) error {
	// get hotel obj and branch object
	branch, err := s.adminService.GetHotelBranch(hotelID, branchID)
	if err != nil {
		return err
	}

	// find available room for room type
	availableRoom, err := branch.FindAvailableRoom(roomType)
	if err != nil {
		return err
	}

	// mark room booked
	err = availableRoom.MarkBooked()
	if err != nil {
		return err
	}

	// create and save booking
	bookingObj := booking.CreateBooking(
		bookingID,
		hotelID,
		availableRoom.GetRoomNumber(),
		branchID,
		guestID,
		duration,
		advance,
	)

	return s.bookingRepo.SaveBooking(bookingObj)
}

func (s *BookingServiceImpl) CancelBooking(bookingID string) error {
	bookingObj, err := s.bookingRepo.GetBooking(bookingID)
	if err != nil {
		return err
	}

	roomObj, err := s.adminService.GetRoom(bookingObj.HotelID, bookingObj.BranchID, bookingObj.RoomID)
	if err != nil {
		return err
	}

	err = roomObj.MarkBooked()
	if err != nil {
		return err
	}

	// Cancel booking and mark room as available
	return s.bookingRepo.CancelBooking(bookingID)
}

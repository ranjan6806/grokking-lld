package booking

import (
	"fmt"
	"hotel-management/hotel_branch"
)

type BookingRepository interface {
	AddBranch(hotelID string, branch *hotel_branch.HotelBranch) error
	RemoveBranch(hotelID, branchID string) error
	GetBranch(hotelID, branchID string) (*hotel_branch.HotelBranch, error)
	SaveBooking(booking *Booking) error
	GetBooking(bookingID string) (*Booking, error)
	CancelBooking(bookingID string) error
}

type BookingRepositoryImpl struct {
	bookings       map[string]*Booking
	activeBranches map[string]map[string]*hotel_branch.HotelBranch
}

func NewBookingRepository() BookingRepository {
	return &BookingRepositoryImpl{
		bookings:       make(map[string]*Booking),
		activeBranches: make(map[string]map[string]*hotel_branch.HotelBranch),
	}
}

func (r *BookingRepositoryImpl) AddBranch(hotelID string, branch *hotel_branch.HotelBranch) error {
	if _, exists := r.activeBranches[hotelID]; !exists {
		r.activeBranches[hotelID] = make(map[string]*hotel_branch.HotelBranch)
	}

	r.activeBranches[hotelID][branch.BranchID] = branch
	return nil
}

func (r *BookingRepositoryImpl) RemoveBranch(hotelID, branchID string) error {
	if _, exists := r.activeBranches[hotelID]; exists {
		if _, exists1 := r.activeBranches[hotelID][branchID]; exists1 {
			delete(r.activeBranches[hotelID], branchID)
			return nil
		}
	}
	return fmt.Errorf("branch %s does not exists", branchID)
}

func (r *BookingRepositoryImpl) GetBranch(hotelID, branchID string) (*hotel_branch.HotelBranch, error) {
	if _, exists := r.activeBranches[hotelID]; exists {
		if _, exists1 := r.activeBranches[hotelID][branchID]; exists1 {
			return r.activeBranches[hotelID][branchID], nil
		}
	}
	return nil, fmt.Errorf("branch %s does not exists", branchID)
}

func (r *BookingRepositoryImpl) SaveBooking(booking *Booking) error {
	if _, exists := r.bookings[booking.BookingID]; exists {
		return fmt.Errorf("booking %s already exists", booking.BookingID)
	}

	r.bookings[booking.BookingID] = booking
	return nil
}

func (r *BookingRepositoryImpl) GetBooking(bookingID string) (*Booking, error) {
	booking, exists := r.bookings[bookingID]
	if !exists {
		return nil, fmt.Errorf("booking %s does not exists", bookingID)
	}

	return booking, nil
}

func (r *BookingRepositoryImpl) CancelBooking(bookingID string) error {
	if _, exists := r.bookings[bookingID]; exists {
		return fmt.Errorf("cannot cancel booking %s", bookingID)
	}

	delete(r.bookings, bookingID)
	return nil
}

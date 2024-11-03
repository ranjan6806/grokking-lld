package booking

import "time"

type BookingStatus string

const (
	BookingStatusCreated   BookingStatus = "CREATED"
	BookingStatusCancelled BookingStatus = "CANCELLED"
)

type Booking struct {
	BookingID      string
	HotelID        string
	BranchID       string
	RoomID         string
	GuestID        string
	DurationInDays uint32
	Status         BookingStatus
	Checkout       time.Time
	AdvancePayment float64
}

func CreateBooking(bookingID, hotelID, roomID, branchID, guestID string, duration uint32, advancePayment float64) *Booking {
	return &Booking{
		BookingID:      bookingID,
		HotelID:        hotelID,
		BranchID:       branchID,
		RoomID:         roomID,
		GuestID:        guestID,
		DurationInDays: duration,
		AdvancePayment: advancePayment,
		Status:         BookingStatusCreated,
	}
}

func (b *Booking) Cancel() error {
	b.Status = BookingStatusCancelled
	return nil
}

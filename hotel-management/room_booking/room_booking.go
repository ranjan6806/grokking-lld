package room_booking

import "time"

type BookingStatus string

const (
	BookingStatusReserved   BookingStatus = "RESERVED"
	BookingStatusCheckedIn  BookingStatus = "CHECKED_IN"
	BookingStatusCheckedOut BookingStatus = "CHECKED_OUT"
)

type RoomBooking struct {
	ReservationNumber string
	StartDate         time.Time
	DurationInDays    int
	Status            BookingStatus
	Checkout          time.Time
	AdvancePayment    float64
}

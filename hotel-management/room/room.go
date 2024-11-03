package room

type RoomType string

const (
	RoomTypeStandard      RoomType = "standard"
	RoomTypeDeluxe        RoomType = "deluxe"
	RoomTypeFamilySuite   RoomType = "family-suite"
	RoomTypeBusinessSuite RoomType = "business-suite"
)

type RoomStatus string

const (
	RoomStatusAvailable     RoomStatus = "available"
	RoomStatusBooked        RoomStatus = "booked"
	RoomStatusOccupied      RoomStatus = "occupied"
	RoomStatusBeingServiced RoomStatus = "being-serviced"
)

type Room interface {
	GetRoomNumber() string
	GetRoomType() RoomType
	IsAvailable() bool
	Checkin() error
	Checkout() error
	MarkBooked() error
}

func NewRoom(roomType RoomType, roomNumber string, bookingPrice float64) Room {
	return &RoomStandard{
		RoomNumber:   roomNumber,
		RoomStatus:   RoomStatusAvailable,
		BookingPrice: bookingPrice,
	}
}

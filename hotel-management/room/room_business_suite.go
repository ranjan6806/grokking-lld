package room

type RoomBusinessSuite struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	BookingPrice float64
}

func (r *RoomBusinessSuite) GetRoomType() RoomType {
	return RoomTypeBusinessSuite
}

func (r *RoomBusinessSuite) MarkBooked() error {
	r.RoomStatus = RoomStatusBooked
	return nil
}

func (r *RoomBusinessSuite) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomBusinessSuite) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomBusinessSuite) Checkin() error {
	return nil
}

func (r *RoomBusinessSuite) Checkout() error {
	return nil
}

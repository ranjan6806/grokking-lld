package room

type RoomDeluxe struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	BookingPrice float64
}

func (r *RoomDeluxe) MarkBooked() error {
	r.RoomStatus = RoomStatusBooked
	return nil
}

func (r *RoomDeluxe) GetRoomType() RoomType {
	return RoomTypeDeluxe
}

func (r *RoomDeluxe) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomDeluxe) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomDeluxe) Checkin() error {
	return nil
}

func (r *RoomDeluxe) Checkout() error {
	return nil
}

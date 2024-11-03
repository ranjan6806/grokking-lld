package room

type RoomStandard struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	BookingPrice float64
}

func (r *RoomStandard) MarkBooked() error {
	r.RoomStatus = RoomStatusBooked
	return nil
}

func (r *RoomStandard) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomStandard) GetRoomType() RoomType {
	return RoomTypeStandard
}

func (r *RoomStandard) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomStandard) Checkin() error {
	return nil
}

func (r *RoomStandard) Checkout() error {
	return nil
}

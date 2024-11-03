package room

type RoomFamilySuite struct {
	RoomNumber   string
	RoomType     RoomType
	RoomStatus   RoomStatus
	BookingPrice float64
}

func (r *RoomFamilySuite) MarkBooked() error {
	r.RoomStatus = RoomStatusBooked
	return nil
}

func (r *RoomFamilySuite) GetRoomType() RoomType {
	return RoomTypeFamilySuite
}

func (r *RoomFamilySuite) GetRoomNumber() string {
	return r.RoomNumber
}

func (r *RoomFamilySuite) IsAvailable() bool {
	return r.RoomStatus == RoomStatusAvailable
}

func (r *RoomFamilySuite) Checkin() error {
	return nil
}

func (r *RoomFamilySuite) Checkout() error {
	return nil
}

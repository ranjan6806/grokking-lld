package room

import "time"

type RoomKey struct {
	KeyId    string
	Barcode  string
	IssuedAt time.Time
	IsActive bool
	IsMaster bool
}

func (rk *RoomKey) AssignRoom(room Room) error {
	return nil
}

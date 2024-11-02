package hotel_branch

import (
	"fmt"
	"hotel-management/address"
	"hotel-management/room"
)

type HotelBranch struct {
	BranchID string
	HotelID  string
	Name     string
	Address  address.Address
	Rooms    []room.Room
}

func NewBranch(id, name string) *HotelBranch {
	return &HotelBranch{
		BranchID: id,
		Rooms:    make([]room.Room, 0),
	}
}

func (h *HotelBranch) AddRoom(room room.Room) error {
	h.Rooms = append(h.Rooms, room)
	return nil
}

func (h *HotelBranch) GetRoom(roomID string) (room.Room, error) {
	for _, room := range h.Rooms {
		if room.GetRoomNumber() == roomID {
			return room, nil
		}
	}
	return nil, fmt.Errorf("room not found")
}

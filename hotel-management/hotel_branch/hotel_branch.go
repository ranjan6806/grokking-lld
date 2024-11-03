package hotel_branch

import (
	"fmt"
	"hotel-management/address"
	"hotel-management/room"
)

type HotelBranch struct {
	BranchID string
	Name     string
	Address  address.Address
	Rooms    map[string]room.Room // map roomID -> Room
}

func NewBranch(id, name string, address address.Address) *HotelBranch {
	return &HotelBranch{
		BranchID: id,
		Name:     name,
		Address:  address,
		Rooms:    make(map[string]room.Room),
	}
}

func (b *HotelBranch) AddRoom(room room.Room) error {
	_, exists := b.Rooms[room.GetRoomNumber()]
	if exists {
		return fmt.Errorf("room %s already exists", room.GetRoomNumber())
	}

	b.Rooms[room.GetRoomNumber()] = room
	return nil
}

func (b *HotelBranch) GetRoom(roomID string) (room.Room, error) {
	for _, r := range b.Rooms {
		if r.GetRoomNumber() == roomID {
			return r, nil
		}
	}
	return nil, fmt.Errorf("room not found")
}

func (b *HotelBranch) FindAvailableRoom(roomType room.RoomType) (room.Room, error) {
	for _, r := range b.Rooms {
		if r.GetRoomType() == roomType && r.IsAvailable() {
			return r, nil
		}
	}

	return nil, fmt.Errorf("room %s does not exists")
}

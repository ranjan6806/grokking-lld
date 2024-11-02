package room

import (
	"errors"
	"sync"
)

type RoomRepository interface {
	AddRoom(room Room) error
	GetRoom(id string) (Room, error)
}

type RoomRepositoryImpl struct {
	rooms map[string]Room // map of room id to Room
	mtx   sync.RWMutex
}

func (r *RoomRepositoryImpl) AddRoom(room Room) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, exists := r.rooms[room.GetRoomNumber()]; exists {
		return errors.New("room already exists")
	}

	r.rooms[room.GetRoomNumber()] = room
	return nil
}

func (r *RoomRepositoryImpl) GetRoom(roomNumber string) (Room, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	room, exists := r.rooms[roomNumber]
	if !exists {
		return nil, errors.New("room not found")
	}
	return room, nil
}

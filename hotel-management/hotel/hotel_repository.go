package hotel

import (
	"fmt"
	"sync"
)

type HotelRepository interface {
	AddHotel(hotel *Hotel) error
	GetHotel(id string) (*Hotel, error)
}

type HotelRepositoryImpl struct {
	hotels map[string]*Hotel
	mtx    sync.RWMutex
}

func NewHotelRepository() HotelRepository {
	return &HotelRepositoryImpl{
		hotels: make(map[string]*Hotel),
	}
}

func (r *HotelRepositoryImpl) AddHotel(hotel *Hotel) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, exists := r.hotels[hotel.HotelID]; exists {
		return fmt.Errorf("hotel with id %s already exists", hotel.HotelID)
	}

	r.hotels[hotel.HotelID] = hotel
	return nil
}

func (r *HotelRepositoryImpl) GetHotel(id string) (*Hotel, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	hotel, exists := r.hotels[id]
	if !exists {
		return nil, fmt.Errorf("hotel with id %s does not exists", id)
	}

	return hotel, nil
}

package service

import "parking-lot/models"

type DisplayBoardServiceInterface interface {
	SetFreeSpots(spotType models.SpotType, count int)
	IncrementFreeSpots(spotType models.SpotType)
	DecrementFreeSpots(spotType models.SpotType)
	ShowFreeSpots() map[models.SpotType]int
}

type DisplayBoardService struct {
	freeSpots map[models.SpotType]int
}

func NewDisplayBoardService() *DisplayBoardService {
	return &DisplayBoardService{
		freeSpots: make(map[models.SpotType]int),
	}
}

func (d *DisplayBoardService) SetFreeSpots(spotType models.SpotType, count int) {
	d.freeSpots[spotType] = count
}

func (d *DisplayBoardService) IncrementFreeSpots(spotType models.SpotType) {
	d.freeSpots[spotType]++
}

func (d *DisplayBoardService) DecrementFreeSpots(spotType models.SpotType) {
	d.freeSpots[spotType]--
}

func (d *DisplayBoardService) ShowFreeSpots() map[models.SpotType]int {
	// create a copy to prevent external modification
	freeSpotsCopy := make(map[models.SpotType]int)
	for k, v := range d.freeSpots {
		freeSpotsCopy[k] = v
	}

	return freeSpotsCopy
}

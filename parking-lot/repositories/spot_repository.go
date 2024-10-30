package repositories

import (
	"errors"
	"parking-lot/models"
)

type SpotRepository struct {
	spots map[int]models.SpotInterface
}

func (sr *SpotRepository) GetAllSpots() ([]models.SpotInterface, error) {
	spots := make([]models.SpotInterface, 0)
	for _, spot := range sr.spots {
		spots = append(spots, spot)
	}
	return spots, nil
}

func (sr *SpotRepository) GetSpotByID(id int) (models.SpotInterface, error) {
	spot, ok := sr.spots[id]
	if !ok {
		return nil, errors.New("Spot not found")
	}
	return spot, nil
}

func (sr *SpotRepository) CreateSpot(spot models.SpotInterface) error {
	if _, exits := sr.spots[spot.GetID()]; exits {
		return errors.New("Spot already exists")
	}

	sr.spots[spot.GetID()] = spot
	return nil
}

func (sr *SpotRepository) UpdateSpot(spot models.SpotInterface) error {
	sr.spots[spot.GetID()] = spot
	return nil
}

func NewSpotRepository() models.SpotRepositoryInterface {
	return &SpotRepository{
		spots: make(map[int]models.SpotInterface),
	}
}

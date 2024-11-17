package services

import (
	"fmt"
	"movie-ticket-booking/models"
)

type AdminServiceInterface interface {
	AddCinema(cinemaID string)
	AddHall(cinemaID string, hallID string)
	AddShow(cinemaID, hallID string, show *models.Show)
}

type AdminService struct {
	cinemas map[string]*models.Cinema
}

func NewAdminService() *AdminService {
	return &AdminService{
		cinemas: make(map[string]*models.Cinema),
	}
}

func (s *AdminService) AddCinema(cinemaID string) {
	s.cinemas[cinemaID] = models.NewCinema(cinemaID)
}

func (s *AdminService) AddHall(cinemaID string, hallID string) {
	cinema, exists := s.cinemas[cinemaID]
	if !exists {
		fmt.Println("cinema does not exist")
		return
	}

	cinema.Halls[hallID] = models.NewHall(hallID)
}

func (s *AdminService) AddShow(cinemaID, hallID string, show *models.Show) {
	cinema, exists := s.cinemas[cinemaID]
	if !exists {
		fmt.Println("cinema does not exist")
		return
	}

	hall, exists := cinema.Halls[hallID]
	if !exists {
		fmt.Println("hall does not exist")
		return
	}

	hall.Shows[show.ShowID] = show
}

func (s *AdminService) ShowAllShows(cinemaID string) {
	cinema, exists := s.cinemas[cinemaID]
	if !exists {
		fmt.Println("cinema does not exist")
		return
	}

	for _, hall := range cinema.Halls {
		for _, show := range hall.Shows {
			fmt.Printf("Show - %+v\n", show)
		}
	}
}

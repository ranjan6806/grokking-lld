package main

import (
	"movie-ticket-booking/models"
	"movie-ticket-booking/services"
	"time"
)

func main() {
	service := services.NewAdminService()
	service.AddCinema("CINEMA1")
	service.AddHall("CINEMA1", "HALL1")
	service.AddShow(
		"CINEMA1",
		"HALL1",
		models.NewShow("SHOW1", models.NewMovie("MOVIE1", "LANGUAGE1", "GENRE1", time.Now())),
	)

	service.ShowAllShows("CINEMA1")
}

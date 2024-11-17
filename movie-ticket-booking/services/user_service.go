package services

import (
	"movie-ticket-booking/models"
	"movie-ticket-booking/search"
)

type UserServiceInterface interface {
	CreateBooking(cinemaID, hallID, showID, bookingID string, user *models.User) *models.Ticket
	CancelBooking(bookingID string)
	SearchShow(queryString string, strategy search.SearchStrategy) []*models.Show
}

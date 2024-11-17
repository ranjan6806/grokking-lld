package models

type Show struct {
	ShowID         string
	Movie          *Movie
	AvailableSeats map[string]*Seat
}

func NewShow(showID string, movie *Movie) *Show {
	return &Show{
		ShowID:         showID,
		Movie:          movie,
		AvailableSeats: make(map[string]*Seat),
	}
}

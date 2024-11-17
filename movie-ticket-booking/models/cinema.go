package models

type Cinema struct {
	CinemaID string
	Halls    map[string]*Hall
}

func NewCinema(cinemaID string) *Cinema {
	return &Cinema{
		CinemaID: cinemaID,
		Halls:    make(map[string]*Hall),
	}
}

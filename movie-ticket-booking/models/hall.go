package models

type Hall struct {
	HallID  string
	SeatMap map[string]*Seat
	Shows   map[string]*Show
}

func NewHall(hallID string) *Hall {
	return &Hall{
		HallID:  hallID,
		SeatMap: make(map[string]*Seat),
		Shows:   make(map[string]*Show),
	}
}

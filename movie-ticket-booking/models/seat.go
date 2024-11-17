package models

const (
	SilverSeatType   = "silver"
	GoldSeatType     = "gold"
	PlatinumSeatType = "platinum"
)

type Seat interface {
	GetID() string
	GetCost() uint
	GetIsBooked() bool
}

type BaseSeat struct {
	ID       string
	Cost     uint
	IsBooked bool
}

func (s *BaseSeat) GetID() string {
	return s.ID
}

func (s *BaseSeat) GetCost() uint {
	return s.Cost
}

func (s *BaseSeat) GetIsBooked() bool {
	return s.IsBooked
}

type SilverSeat struct {
	BaseSeat
}

type GoldSeat struct {
	BaseSeat
}

type PlatinumSeat struct {
	BaseSeat
}

func NewSeat(id string, cost uint, seatType string) Seat {
	switch seatType {
	case "silver":
		return &SilverSeat{
			BaseSeat: BaseSeat{
				ID:   id,
				Cost: cost,
			},
		}
	case "gold":
		return &GoldSeat{
			BaseSeat: BaseSeat{
				ID:   id,
				Cost: cost,
			},
		}
	case "platinum":
		return &PlatinumSeat{
			BaseSeat: BaseSeat{
				ID:   id,
				Cost: cost,
			},
		}
	default:
		return nil
	}
}

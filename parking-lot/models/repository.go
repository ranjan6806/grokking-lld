package models

type SpotRepositoryInterface interface {
	GetAllSpots() ([]SpotInterface, error)
	GetSpotByID(spotID int) (SpotInterface, error)
	UpdateSpot(spot SpotInterface) error
	CreateSpot(spot SpotInterface) error
}

type TicketRepositoryInterface interface {
	GetAllTickets() ([]TicketInterface, error)
	GetTicketByID(ticketID string) (TicketInterface, error)
	CreateTicket(ticket TicketInterface) error
	DeleteTicket(ticketID string) error
}

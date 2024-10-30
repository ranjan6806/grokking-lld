package repositories

import (
	"errors"
	"parking-lot/models"
)

type TicketRepository struct {
	tickets map[string]models.TicketInterface
}

func (tr *TicketRepository) GetAllTickets() ([]models.TicketInterface, error) {
	tickets := []models.TicketInterface{}
	for _, ticket := range tr.tickets {
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (tr *TicketRepository) GetTicketByID(ticketId string) (models.TicketInterface, error) {
	ticket, ok := tr.tickets[ticketId]
	if !ok {
		return nil, errors.New("Ticket not found")
	}
	return ticket, nil
}

func (tr *TicketRepository) CreateTicket(ticket models.TicketInterface) error {
	if _, exists := tr.tickets[ticket.GetID()]; exists {
		return errors.New("Ticket already exists")
	}
	tr.tickets[ticket.GetID()] = ticket
	return nil
}

func (tr *TicketRepository) DeleteTicket(ticketId string) error {
	if _, exists := tr.tickets[ticketId]; exists {
		delete(tr.tickets, ticketId)
		return nil
	}
	return errors.New("Ticket not found")
}

// NewTicketRepository creates a new TicketRepository instance.
func NewTicketRepository() models.TicketRepositoryInterface {
	return &TicketRepository{
		tickets: make(map[string]models.TicketInterface),
	}
}

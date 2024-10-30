package models

type PaymentMethod string

const (
	CreditCard PaymentMethod = "CreditCard"
	DebitCard  PaymentMethod = "DebitCard"
	Cash       PaymentMethod = "Cash"
)

type PaymentInterface interface {
	GetAmount() float64
	GetPaymentMethod() PaymentMethod
	GetTicket() TicketInterface
}

type Payment struct {
	Ticket        TicketInterface
	Amount        float64
	PaymentMethod PaymentMethod
}

func (p *Payment) GetAmount() float64 {
	return p.Amount
}

func (p *Payment) GetPaymentMethod() PaymentMethod {
	return p.PaymentMethod
}

func (p *Payment) GetTicket() TicketInterface {
	return p.Ticket
}

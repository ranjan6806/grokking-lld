package service

import "parking-lot/models"

type PaymentServiceInterface interface {
	ProcessPayment(ticket models.TicketInterface, amount float64, method models.PaymentMethod) models.PaymentInterface
}

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (ps *PaymentService) ProcessPayment(ticket models.TicketInterface, amount float64, method models.PaymentMethod) models.PaymentInterface {
	return &models.Payment{
		Ticket:        ticket,
		Amount:        amount,
		PaymentMethod: method,
	}
}

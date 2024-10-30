package service

import (
	"errors"
	"parking-lot/models"
)

type PaymentServiceInterface interface {
	ProcessPayment(payment models.PaymentInterface, amount float64) (bool, error)
}

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (ps *PaymentService) ProcessPayment(payment models.PaymentInterface, amount float64) (bool, error) {
	if payment == nil {
		return false, errors.New("payment is nil")
	}

	return payment.ProcessPayment(amount)
}

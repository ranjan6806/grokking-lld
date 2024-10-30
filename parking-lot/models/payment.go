package models

import "errors"

type PaymentInterface interface {
	ProcessPayment(amount float64) (bool, error)
	GetPaymentType() PaymentType
}

type CashPayment struct{}

func (c *CashPayment) ProcessPayment(amount float64) (bool, error) {
	if amount <= 0 {
		return false, errors.New("invalid amount")
	}

	// Simulate cash payment processing
	return true, nil
}

func (c *CashPayment) GetPaymentType() PaymentType {
	return Cash
}

type CreditCardPayment struct {
	CardNumber string
	Expiry     string
	CVV        string
}

func (cc *CreditCardPayment) ProcessPayment(amount float64) (bool, error) {
	if amount <= 0 {
		return false, errors.New("invalid amount")
	}

	if len(cc.CardNumber) < 12 {
		return false, errors.New("invalid card number")
	}

	// Simulate credit card payment processing
	return true, nil
}

func (cc *CreditCardPayment) GetPaymentType() PaymentType {
	return CreditCard
}

func NewPayment(paymentType PaymentType, details ...string) (PaymentInterface, error) {
	switch paymentType {
	case Cash:
		return &CashPayment{}, nil
	case CreditCard:
		if len(details) < 3 {
			return nil, errors.New("invalid credit card")
		}
		return &CreditCardPayment{
			CardNumber: details[0],
			Expiry:     details[1],
			CVV:        details[2],
		}, nil
	default:
		return nil, errors.New("invalid payment type")
	}
}

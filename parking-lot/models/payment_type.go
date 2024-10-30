package models

type PaymentType string

const (
	CreditCard PaymentType = "CreditCard"
	DebitCard  PaymentType = "DebitCard"
	Cash       PaymentType = "Cash"
)

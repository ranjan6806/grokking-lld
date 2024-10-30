package models

import "time"

type TicketInterface interface {
	GetID() string
	GetVehicle() VehicleInterface
	GetSpot() SpotInterface
	GetEntryTime() time.Time
	GetPayment() PaymentInterface
}

type Ticket struct {
	ID        string
	Vehicle   VehicleInterface
	Spot      SpotInterface
	EntryTime time.Time
	Payment   PaymentInterface
}

func (t *Ticket) GetID() string {
	return t.ID
}

func (t *Ticket) GetVehicle() VehicleInterface {
	return t.Vehicle
}

func (t *Ticket) GetSpot() SpotInterface {
	return t.Spot
}

func (t *Ticket) GetEntryTime() time.Time {
	return t.EntryTime
}

func (t *Ticket) GetPayment() PaymentInterface {
	return t.Payment
}

func NewTicket(vehicle VehicleInterface, spot SpotInterface, payment PaymentInterface) TicketInterface {
	return &Ticket{
		ID:        vehicle.GetLicensePlate(),
		Vehicle:   vehicle,
		Spot:      spot,
		EntryTime: time.Now(),
		Payment:   payment,
	}
}

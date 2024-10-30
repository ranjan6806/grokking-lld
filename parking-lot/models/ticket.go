package models

import "time"

type TicketInterface interface {
	GetID() string
	GetVehicle() VehicleInterface
	GetSpot() SpotInterface
	GetEntryTime() time.Time
}

type Ticket struct {
	ID        string
	Vehicle   VehicleInterface
	Spot      SpotInterface
	EntryTime time.Time
}

func NewTicket(vehicle VehicleInterface, spot SpotInterface) TicketInterface {
	return &Ticket{
		ID:        vehicle.GetLicensePlate(),
		Vehicle:   vehicle,
		Spot:      spot,
		EntryTime: time.Now(),
	}
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

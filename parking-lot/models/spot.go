package models

type SpotType string

const (
	Handicapped    SpotType = "Handicapped"
	Compact        SpotType = "Compact"
	Large          SpotType = "Large"
	MotorcycleSpot SpotType = "MotorcycleSpot"
)

type SpotInterface interface {
	GetID() int
	GetSpotType() SpotType
	IsFree() bool
	ParkVehicle(VehicleInterface)
	RemoveVehicle()
}

type Spot struct {
	ID       int
	SpotType SpotType
	Free     bool
	Vehicle  VehicleInterface
}

func (s *Spot) GetID() int {
	return s.ID
}

func (s *Spot) GetSpotType() SpotType {
	return s.SpotType
}

func (s *Spot) IsFree() bool {
	return s.Free
}

func (s *Spot) ParkVehicle(vehicle VehicleInterface) {
	s.Vehicle = vehicle
	s.Free = true
}

func (s *Spot) RemoveVehicle() {
	s.Vehicle = nil
	s.Free = true
}

// Factory function returns SpotInterface
func NewSpot(id int, spotType SpotType) SpotInterface {
	return &Spot{ID: id, SpotType: spotType, Free: true}
}

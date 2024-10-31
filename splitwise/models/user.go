package models

type User struct {
	ID      string
	Name    string
	Balance map[string]float64
}

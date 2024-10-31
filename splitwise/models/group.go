package models

type Group struct {
	ID      string
	Name    string
	Members []*User
}

package models

type UserInterface interface {
	GetID() string
	GetName() string
	GetRole() Role
}

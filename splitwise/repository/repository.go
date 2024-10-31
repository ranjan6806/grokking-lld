package repository

import "splitwise/models"

type UserRepositoryInterface interface {
	AddUser(user *models.User)
	RemoveUser(userID string)
	GetUser(userID string) *models.User
}

type GroupRepositoryInterface interface {
	AddGroup(group *models.Group)
	RemoveGroup(groupID string)
	GetGroup(groupID string) *models.Group
}

type BalanceRepositoryInterface interface {
	SaveBalance(userID string, balance map[string]float64)
	GetBalance(userID string) map[string]float64
	UpdateBalance(userID string, balance map[string]float64)
	ClearBalances()
	GetAllUsers() []string
}

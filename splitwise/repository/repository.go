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
	// User level balance management
	SaveBalance(userID string, balance map[string]float64)
	GetBalance(userID string) map[string]float64
	UpdateBalance(userID string, balance map[string]float64)
	ClearBalances()
	GetAllUsers() []string

	// Group level balance management
	SaveGroupBalance(groupID string, balance map[string]map[string]float64)
	GetGroupBalance(groupID, userID string) map[string]float64
	UpdateGroupBalance(groupID string, userID string, balance map[string]float64)
	GetAllGroupUsers(groupID string) []string
}

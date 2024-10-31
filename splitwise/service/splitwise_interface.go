package service

import (
	"splitwise/expenses"
	"splitwise/models"
)

type SplitwiseServiceInterface interface {
	AddUser(user *models.User)
	CreateGroup(group *models.Group)
	AddExpense(amount float64, payer *models.User, strategy expenses.SplitStrategyInterface, receivers []*models.User)
	Settle(user1ID, user2ID string)
	ShowAllBalances() map[string]map[string]float64
}

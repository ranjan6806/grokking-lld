package balance

import "splitwise/models"

type BalanceManagerInterface interface {
	// User level balance management
	AddExpense(expense map[string]float64, payer *models.User)
	AddUserBalance(user *models.User)
	ShowBalances() map[string]map[string]float64
	SettleBetweenUsers(user1ID, user2ID string)

	// Group level balance management
	AddGroupExpense(groupID string, expense map[string]float64, payer *models.User, members []*models.User)
	SettleGroup(groupID string)
}

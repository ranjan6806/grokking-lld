package balance

import "splitwise/models"

type BalanceManagerInterface interface {
	AddExpense(expense map[string]float64, payer *models.User)
	AddUserBalance(user *models.User)
	ShowBalances() map[string]map[string]float64
	SettleBetweenUsers(user1ID, user2ID string)
}

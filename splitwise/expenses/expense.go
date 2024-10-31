package expenses

import "splitwise/models"

type Expense struct {
	Amount    float64
	Payer     *models.User
	Strategy  SplitStrategyInterface
	Receivers []*models.User
}

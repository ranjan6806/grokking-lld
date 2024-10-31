package expenses

import "splitwise/models"

type ExpenseFactory struct{}

func (f *ExpenseFactory) CreateExpense(
	amount float64,
	payer *models.User,
	strategy SplitStrategyInterface,
	receivers []*models.User,
) *Expense {
	return &Expense{
		Amount:    amount,
		Payer:     payer,
		Strategy:  strategy,
		Receivers: receivers,
	}
}

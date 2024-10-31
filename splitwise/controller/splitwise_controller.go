package controller

import (
	"fmt"
	"splitwise/expenses"
	"splitwise/models"
	"splitwise/service"
)

type SplitwiseController struct {
	Service service.SplitwiseServiceInterface
}

func NewSplitwiseController(service service.SplitwiseServiceInterface) *SplitwiseController {
	return &SplitwiseController{
		Service: service,
	}
}

func (c *SplitwiseController) AddUser(name string) {
	user := &models.User{Name: name}
	c.Service.AddUser(user)
	fmt.Printf("User %s added with ID %s\n", name, user.ID)
}

func (c *SplitwiseController) CreateGroup(name string, members []*models.User) {
	group := &models.Group{Name: name, Members: members}
	c.Service.CreateGroup(group)
	fmt.Printf("Group %s created with ID %s\n", group.Name, group.ID)
}

func (c *SplitwiseController) AddExpense(
	amount float64,
	payer *models.User,
	strategy expenses.SplitStrategyInterface,
	receivers []*models.User,
) {
	c.Service.AddExpense(amount, payer, strategy, receivers)
	fmt.Println("Expense added")
}

func (c *SplitwiseController) Settle(user1ID, user2ID string) {
	c.Service.Settle(user1ID, user2ID)
	fmt.Printf("Settlement done between %s and %s\n", user1ID, user2ID)
}

func (c *SplitwiseController) ShowAllBalances() {
	balances := c.Service.ShowAllBalances()
	fmt.Printf("All Balances\n")
	for userID, balance := range balances {
		fmt.Printf("User %s balances: %v\n", userID, balance)
	}
}

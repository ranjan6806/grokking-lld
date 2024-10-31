package main

import (
	"fmt"
	"splitwise/balance"
	"splitwise/controller"
	"splitwise/expenses"
	"splitwise/models"
	"splitwise/repository"
	"splitwise/service"
)

func main() {
	// Initialize repositories, balance manager, service and controller
	userRepo := repository.NewInMemoryUserRepository()
	groupRepo := repository.NewGroupRepository()
	balanceRepo := repository.GetBalanceRepositoryInstance()

	// Initialize balance manager with repository
	balanceManager := balance.NewBalanceManager(balanceRepo)
	splitwiseService := service.NewSplitwiseService(userRepo, groupRepo, balanceManager)
	splitwiseController := controller.NewSplitwiseController(splitwiseService)

	splitwiseController.AddUser("Alice")
	splitwiseController.AddUser("Bob")

	user1 := userRepo.GetUser("u1")
	user2 := userRepo.GetUser("u2")

	splitwiseController.AddExpense(400, user1, &expenses.EqualSplit{}, []*models.User{user1, user2})

	//fmt.Println("All Balances")
	splitwiseController.ShowAllBalances()

	splitwiseController.AddExpense(800, user2, &expenses.EqualSplit{}, []*models.User{user1, user2})

	splitwiseController.ShowAllBalances()

	splitwiseController.Settle("u1", "u2")
	fmt.Println("Balances after settlement")
	splitwiseController.ShowAllBalances()

	// Define group members
	user1 = &models.User{ID: "u1", Name: "Alice"}
	user2 = &models.User{ID: "u2", Name: "Bob"}
	user3 := &models.User{ID: "u3", Name: "Charlie"}

	members := []*models.User{user1, user2, user3}
	groupID := "g1"

	// Initialize group balance
	initialBalances := map[string]map[string]float64{
		"u1": {"u2": 0, "u3": 0},
		"u2": {"u1": 0, "u3": 0},
		"u3": {"u1": 0, "u2": 0},
	}
	balanceRepo.SaveGroupBalance(groupID, initialBalances)

	// Add a group expense where Alice paid $150, split equally
	expense := map[string]float64{user1.ID: 150.0}
	balanceManager.AddGroupExpense(groupID, expense, user1, members)

	// Show group balances (each of Bob and Charlie owe Alice $50)
	groupBalances := balanceRepo.GetGroupBalance(groupID, user1.ID)
	fmt.Printf("Group Balances for Alice: %+v\n", groupBalances)

	// Settle the group balances
	balanceManager.SettleGroup(groupID)
	groupBalances = balanceRepo.GetGroupBalance(groupID, user1.ID)
	fmt.Printf("Group Balances for Alice after settlement: %+v\n", groupBalances)
}

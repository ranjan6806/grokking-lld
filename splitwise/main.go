package main

import (
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
	balanceRepo := repository.NewInMemoryBalanceRepository()

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

	//splitwiseController.Settle("u1", "u2")
	//fmt.Println("Balances after settlement")
	//splitwiseController.ShowAllBalances()
}

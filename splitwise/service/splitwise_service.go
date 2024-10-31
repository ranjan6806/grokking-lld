package service

import (
	"splitwise/balance"
	"splitwise/expenses"
	"splitwise/models"
	"splitwise/repository"
)

type SplitwiseService struct {
	UserRepo       repository.UserRepositoryInterface
	GroupRepo      repository.GroupRepositoryInterface
	BalanceManager balance.BalanceManagerInterface
	ExpenseFactory *expenses.ExpenseFactory
}

func (s *SplitwiseService) AddUser(user *models.User) {
	s.UserRepo.AddUser(user)
	s.BalanceManager.AddUserBalance(user)
}

func (s *SplitwiseService) CreateGroup(group *models.Group) {
	s.GroupRepo.AddGroup(group)
}

func (s *SplitwiseService) AddExpense(
	amount float64,
	payer *models.User,
	strategy expenses.SplitStrategyInterface,
	receivers []*models.User,
) {
	expense := s.ExpenseFactory.CreateExpense(amount, payer, strategy, receivers)
	splits := expense.Strategy.CalculateSplit(expense.Amount, expense.Receivers)
	s.BalanceManager.AddExpense(splits, payer)
}

func (s *SplitwiseService) Settle(user1ID, user2ID string) {
	s.BalanceManager.SettleBetweenUsers(user1ID, user2ID)
}

func (s *SplitwiseService) ShowAllBalances() map[string]map[string]float64 {
	return s.BalanceManager.ShowBalances()
}

func NewSplitwiseService(
	userRepo repository.UserRepositoryInterface,
	groupRepo repository.GroupRepositoryInterface,
	balanceManager balance.BalanceManagerInterface,
) SplitwiseServiceInterface {
	return &SplitwiseService{
		UserRepo:       userRepo,
		GroupRepo:      groupRepo,
		BalanceManager: balanceManager,
		ExpenseFactory: &expenses.ExpenseFactory{},
	}
}

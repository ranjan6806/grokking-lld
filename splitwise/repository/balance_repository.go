package repository

type InMemoryBalanceRepository struct {
	balances map[string]map[string]float64 // Map of userID -> balances with other users
}

func (br *InMemoryBalanceRepository) SaveBalance(userID string, balance map[string]float64) {
	br.balances[userID] = balance
}

func (br *InMemoryBalanceRepository) GetBalance(userID string) map[string]float64 {
	return br.balances[userID]
}

func (br *InMemoryBalanceRepository) UpdateBalance(userID string, balance map[string]float64) {
	br.balances[userID] = balance
}

func (br *InMemoryBalanceRepository) ClearBalances() {
	br.balances = make(map[string]map[string]float64)
}

func (br *InMemoryBalanceRepository) GetAllUsers() []string {
	userIDs := make([]string, 0)
	for userID := range br.balances {
		userIDs = append(userIDs, userID)
	}
	return userIDs
}

func NewInMemoryBalanceRepository() BalanceRepositoryInterface {
	return &InMemoryBalanceRepository{
		balances: make(map[string]map[string]float64),
	}
}

package repository

import "sync"

var (
	instance *InMemoryBalanceRepository
	once     sync.Once
)

type InMemoryBalanceRepository struct {
	balances      map[string]map[string]float64            // Map of userID -> balances with other users
	groupBalances map[string]map[string]map[string]float64 // Group level balances: groupID -> (userID -> (otherUserID -> amount))

	userMutes  sync.RWMutex
	groupMutes sync.RWMutex
}

func (br *InMemoryBalanceRepository) SaveBalance(userID string, balance map[string]float64) {
	br.userMutes.Lock()
	defer br.userMutes.Unlock()

	br.balances[userID] = balance
}

func (br *InMemoryBalanceRepository) GetBalance(userID string) map[string]float64 {
	br.userMutes.RLock()
	defer br.userMutes.RUnlock()

	return br.balances[userID]
}

func (br *InMemoryBalanceRepository) UpdateBalance(userID string, balance map[string]float64) {
	br.userMutes.Lock()
	defer br.userMutes.Unlock()

	br.balances[userID] = balance
}

func (br *InMemoryBalanceRepository) ClearBalances() {
	br.userMutes.Lock()
	defer br.userMutes.Unlock()

	br.balances = make(map[string]map[string]float64)
}

func (br *InMemoryBalanceRepository) GetAllUsers() []string {
	br.userMutes.RLock()
	defer br.userMutes.RUnlock()

	userIDs := make([]string, 0)
	for userID := range br.balances {
		userIDs = append(userIDs, userID)
	}
	return userIDs
}

func (br *InMemoryBalanceRepository) SaveGroupBalance(groupID string, balance map[string]map[string]float64) {
	br.groupMutes.Lock()
	defer br.groupMutes.Unlock()

	br.groupBalances[groupID] = balance
}

func (br *InMemoryBalanceRepository) GetGroupBalance(groupID, userID string) map[string]float64 {
	br.groupMutes.RLock()
	defer br.groupMutes.RUnlock()

	if group, ok := br.groupBalances[groupID]; ok {
		if balance, ok := group[userID]; ok {
			return balance
		}
	}

	return make(map[string]float64)
}

func (br *InMemoryBalanceRepository) UpdateGroupBalance(groupID, userID string, balance map[string]float64) {
	br.groupMutes.Lock()
	defer br.groupMutes.Unlock()

	if _, ok := br.groupBalances[groupID]; ok {
		if _, ok := br.groupBalances[groupID][userID]; ok {
			br.groupBalances[groupID][userID] = balance
		}
	}
}

func (br *InMemoryBalanceRepository) GetAllGroupUsers(groupID string) []string {
	br.groupMutes.RLock()
	defer br.groupMutes.RUnlock()

	groupUsers, ok := br.groupBalances[groupID]
	if !ok {
		return make([]string, 0)
	}

	userIDs := make([]string, 0)
	for userID := range groupUsers {
		userIDs = append(userIDs, userID)
	}

	return userIDs
}

func GetBalanceRepositoryInstance() BalanceRepositoryInterface {
	once.Do(func() {
		instance = &InMemoryBalanceRepository{
			balances:      make(map[string]map[string]float64),
			groupBalances: make(map[string]map[string]map[string]float64),
		}
	})

	return instance
}

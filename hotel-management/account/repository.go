package account

import (
	"fmt"
	"sync"
)

type AccountRepository interface {
	AddAccount(account Account) error
	GetAccount(accountID string) (Account, error)
}

type AccountRepositoryImpl struct {
	accounts map[string]*Account // map of account id to accounts
	mtx      sync.RWMutex
}

func (r *AccountRepositoryImpl) AddAccount(account Account) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, exists := r.accounts[account.GetID()]; exists {
		return fmt.Errorf("account with id %s already exists", account.GetID())
	}

	r.accounts[account.GetID()] = &account
	return nil
}

func (r *AccountRepositoryImpl) GetAccount(accountID string) (Account, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	account, exists := r.accounts[accountID]
	if !exists {
		return nil, fmt.Errorf("account with id %s does not exists", accountID)
	}

	return *account, nil
}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{
		accounts: make(map[string]*Account),
	}
}

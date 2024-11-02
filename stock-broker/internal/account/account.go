package account

import (
	"stock-broker/internal/watchlist"
	"stock-broker/repository/watchlist_repository"
)

type Account struct {
	ID                  string
	Balance             float64
	watchlistRepository watchlist_repository.WatchlistRepository
}

func NewAccount(id string, watchListRepo watchlist_repository.WatchlistRepository) *Account {
	return &Account{
		ID:                  id,
		Balance:             0.0,
		watchlistRepository: watchListRepo,
	}
}

func (a *Account) GetID() string {
	return a.ID
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) AddWatchList(watchList *watchlist.WatchList) error {
	return a.watchlistRepository.AddWatchList(a.ID, watchList)
}

func (a *Account) GetWatchList(watchListID string) (*watchlist.WatchList, error) {
	return a.watchlistRepository.GetWatchList(a.ID, watchListID)
}

func (a *Account) ListWatchLists() ([]*watchlist.WatchList, error) {
	return a.watchlistRepository.ListAccountWatchLists(a.ID)
}

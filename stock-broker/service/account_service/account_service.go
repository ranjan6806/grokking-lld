package account_service

import (
	"stock-broker/internal/watchlist"
)

type AccountService interface {
	AddWatchList(accountID string, list *watchlist.WatchList) error
	RemoveWatchList(accountID string, watchListID string) error
	GetWatchLists(accountID string) (watchlist.WatchList, error)
}

type AccountServiceImpl struct{}

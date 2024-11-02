package watchlist_repository

import (
	"errors"
	"stock-broker/internal/watchlist"
)

type WatchlistRepository interface {
	AddWatchList(accountID string, list *watchlist.WatchList) error
	GetWatchList(accountID, watchListId string) (*watchlist.WatchList, error)
	RemoveWatchList(accountID, watchListID string) error
	ListAccountWatchLists(accountID string) ([]*watchlist.WatchList, error)
}

type InMemoryWatchlistRepository struct {
	watchLists map[string]map[string]*watchlist.WatchList // map[accountID][watchlistID]*Watchlist
}

func (wlr *InMemoryWatchlistRepository) AddWatchList(accountID string, watchList *watchlist.WatchList) error {
	if _, exists := wlr.watchLists[accountID]; !exists {
		wlr.watchLists[accountID] = make(map[string]*watchlist.WatchList)
	}

	if _, exists := wlr.watchLists[accountID][watchList.Id]; exists {
		return errors.New("watch list already exists")
	}

	wlr.watchLists[accountID][watchList.Id] = watchList
	return nil
}

func (wlr *InMemoryWatchlistRepository) GetWatchList(accountID, watchListID string) (*watchlist.WatchList, error) {
	if accountWatchLists, exists := wlr.watchLists[accountID]; exists {
		return accountWatchLists[watchListID], nil
	}

	return nil, errors.New("account watch list does not exist")
}

func (wlr *InMemoryWatchlistRepository) RemoveWatchList(accountID, watchListID string) error {
	if accountWatchLists, exists := wlr.watchLists[accountID]; exists {
		delete(accountWatchLists, watchListID)
	}
	return errors.New("watch list does not exist")
}

func (wlr *InMemoryWatchlistRepository) ListAccountWatchLists(accountID string) ([]*watchlist.WatchList, error) {
	if accountWatchLists, exists := wlr.watchLists[accountID]; exists {
		result := make([]*watchlist.WatchList, 0)
		for _, watchList := range accountWatchLists {
			result = append(result, watchList)
		}
		return result, nil
	}

	return nil, errors.New("account watch list does not exist")
}

func NewInMemoryWatchlistRepository() *InMemoryWatchlistRepository {
	return &InMemoryWatchlistRepository{watchLists: make(map[string]map[string]*watchlist.WatchList)}
}

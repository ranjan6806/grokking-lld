package watchlist

import (
	"fmt"
	"stock-broker/internal/stock"
)

type WatchList struct {
	Id     string
	Name   string
	Stocks map[string]*stock.Stock
}

func NewWatchList(id, name string) *WatchList {
	return &WatchList{
		Id:     id,
		Name:   name,
		Stocks: make(map[string]*stock.Stock),
	}
}

func (w *WatchList) AddStock(stock *stock.Stock) error {
	if _, exists := w.Stocks[stock.Name]; exists {
		return fmt.Errorf("stock %s already exists", stock.Name)
	}
	w.Stocks[stock.Ticker] = stock
	return nil
}

func (w *WatchList) RemoveStock(ticker string) error {
	if _, exists := w.Stocks[ticker]; !exists {
		return fmt.Errorf("stock %s does not exists", ticker)
	}
	delete(w.Stocks, ticker)
	return nil
}

func (w *WatchList) GetStocks() []*stock.Stock {
	stocks := make([]*stock.Stock, 0)
	for _, stock := range w.Stocks {
		stocks = append(stocks, stock)
	}
	return stocks
}

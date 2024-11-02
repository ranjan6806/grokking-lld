package stock_repository

import (
	"errors"
	stock2 "stock-broker/internal/stock"
)

type StockRepository interface {
	AddStock(stock *stock2.Stock) error
	GetStock(ticker string) (*stock2.Stock, error)
}

type InMemoryStockRepository struct {
	stocks map[string]*stock2.Stock
}

func (sr *InMemoryStockRepository) AddStock(stock *stock2.Stock) error {
	if _, exists := sr.stocks[stock.Ticker]; exists {
		return errors.New("stock already exists")
	}

	sr.stocks[stock.Ticker] = stock
	return nil
}

func (sr *InMemoryStockRepository) GetStock(ticker string) (*stock2.Stock, error) {
	if stock, exists := sr.stocks[ticker]; exists {
		return stock, nil
	}

	return nil, errors.New("stock not found")
}

func NewInMemoryStockRepository() StockRepository {
	return &InMemoryStockRepository{stocks: make(map[string]*stock2.Stock)}
}

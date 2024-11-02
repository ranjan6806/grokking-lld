package portfolio

import (
	"stock-broker/internal/stock"
	"stock-broker/internal/stock_lot"
)

type Portfolio struct {
	stockLots map[string][]*stock_lot.StockLot
}

func NewPortfolio() *Portfolio {
	return &Portfolio{stockLots: make(map[string][]*stock_lot.StockLot)}
}

func (p *Portfolio) AddStockLot(lot *stock_lot.StockLot) {
	p.stockLots[lot.Stock.Ticker] = append(p.stockLots[lot.Stock.Ticker], lot)
}

type Watchlist struct {
	Stocks map[string]*stock.Stock
}

func NewWatchlist() *Watchlist {
	return &Watchlist{Stocks: make(map[string]*stock.Stock)}
}

func (w *Watchlist) AddStock(stock *stock.Stock) {
	w.Stocks[stock.Ticker] = stock
}

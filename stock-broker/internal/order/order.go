package order

import (
	"fmt"
	"stock-broker/internal/stock"
)

type OrderType string

const (
	MarketOrderType OrderType = "MARKET"
	LimitOrderType  OrderType = "LIMIT"
)

type Order interface {
	Execute() string
}

type MarketOrder struct {
	Stock    *stock.Stock
	Quantity int
}

func (m *MarketOrder) Execute() string {
	return fmt.Sprintf("Executing market order for %d shares of %s", m.Quantity, m.Stock.Ticker)
}

type LimitOrder struct {
	Stock      *stock.Stock
	Quantity   int
	LimitPrice float64
}

func (l *LimitOrder) Execute() string {
	return fmt.Sprintf("Execute limit order for %d shares of %s at %2.2f", l.Quantity, l.Stock.Ticker, l.LimitPrice)
}

type OrderFactory struct{}

func (f *OrderFactory) CreateOrder(
	orderType OrderType,
	stock *stock.Stock,
	quantity int,
	limitPrice float64,
) Order {
	switch orderType {
	case MarketOrderType:
		return &MarketOrder{Stock: stock, Quantity: quantity}
	case LimitOrderType:
		return &LimitOrder{Stock: stock, Quantity: quantity, LimitPrice: limitPrice}
	default:
		return nil
	}
}

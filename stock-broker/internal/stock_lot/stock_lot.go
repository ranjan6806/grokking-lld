package stock_lot

import (
	"stock-broker/internal/stock"
)

type StockLot struct {
	ID            string
	Stock         stock.Stock
	Quantity      int
	PurchasePrice float64
}

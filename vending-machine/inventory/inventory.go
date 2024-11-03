package inventory

import (
	"fmt"
	"sync"
	"vending-machine/product"
)

type Inventory interface {
	AddProduct(row, column uint, product2 *product.Product) error
	RemoveProduct(row, column uint) error
	GetProduct(row, column uint) (*product.Product, error)
	GetAllProducts() []*product.Product
}

type InventoryImpl struct {
	maxRows        uint
	maxColumns     uint
	storedProducts map[uint]map[uint]*product.Product // row -> column -> product
	mtx            sync.RWMutex
}

var (
	inventoryInstance *InventoryImpl
	once              sync.Once
)

func NewInventory(maxRows, maxColumns uint) Inventory {
	once.Do(func() {
		inventoryInstance = &InventoryImpl{
			maxRows:        maxRows,
			maxColumns:     maxColumns,
			storedProducts: make(map[uint]map[uint]*product.Product),
		}
	})
	return inventoryInstance
}

func (i *InventoryImpl) AddProduct(row, column uint, prod *product.Product) error {
	i.mtx.Lock()
	defer i.mtx.Unlock()
	if row > i.maxRows || column > i.maxColumns {
		return fmt.Errorf("invalid row and column")
	}

	if _, ok := i.storedProducts[row]; !ok {
		i.storedProducts[row] = make(map[uint]*product.Product)
	}

	if _, exists := i.storedProducts[row][column]; exists {
		return fmt.Errorf("product already exists")
	}

	i.storedProducts[row][column] = prod
	return nil
}

func (i *InventoryImpl) RemoveProduct(row, column uint) error {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	if row > i.maxRows || column > i.maxColumns {
		return fmt.Errorf("invalid row and column")
	}

	if _, ok := i.storedProducts[row]; !ok {
		return fmt.Errorf("product does not exist")
	}

	if _, ok := i.storedProducts[row][column]; !ok {
		return fmt.Errorf("product does not exist")
	}

	delete(i.storedProducts[row], column)
	return nil
}

func (i *InventoryImpl) GetProduct(row, column uint) (*product.Product, error) {
	i.mtx.RLock()
	defer i.mtx.RUnlock()
	if row > i.maxRows || column > i.maxColumns {
		return nil, fmt.Errorf("invalid row and column")
	}

	if _, ok := i.storedProducts[row]; !ok {
		return nil, fmt.Errorf("product does not exist")
	}

	if _, ok := i.storedProducts[row][column]; !ok {
		return nil, fmt.Errorf("product does not exist")
	}

	return i.storedProducts[row][column], nil
}

func (i *InventoryImpl) GetAllProducts() []*product.Product {
	i.mtx.RLock()
	defer i.mtx.RUnlock()
	products := make([]*product.Product, 0)

	for _, rowProducts := range i.storedProducts {
		for _, prd := range rowProducts {
			products = append(products, prd)
		}
	}

	return products
}

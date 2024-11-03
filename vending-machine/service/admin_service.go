package service

import (
	"fmt"
	"vending-machine/machine"
	"vending-machine/product"
)

type AdminService struct {
	vendingMachine *machine.VendingMachine
}

func NewAdminService(vendingMachine *machine.VendingMachine) *AdminService {
	return &AdminService{
		vendingMachine: vendingMachine,
	}
}

func (a *AdminService) AddProduct(row, column uint, product *product.Product) error {
	return a.vendingMachine.AddProduct(row, column, product)
}

func (a *AdminService) RemoveProduct(row, column uint) error {
	return a.vendingMachine.RemoveProduct(row, column)
}

func (a *AdminService) ShowAllProducts() {
	allProducts := a.vendingMachine.GetAllProducts()

	fmt.Println("ALL PRODUCTS")
	for _, product := range allProducts {
		fmt.Printf("Product - %+v\n", product)
	}
}

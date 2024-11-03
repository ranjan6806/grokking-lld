package machine

import (
	"vending-machine/inventory"
	"vending-machine/product"
)

type VendingMachine struct {
	CurrentState  State
	Inventory     inventory.Inventory
	CurrentAmount uint
}

func NewVendingMachine(inventory inventory.Inventory) *VendingMachine {
	vendingMachine := &VendingMachine{
		Inventory:     inventory,
		CurrentAmount: 0,
	}

	vendingMachine.CurrentState = NewNoMoneyState(vendingMachine)
	return vendingMachine
}

func (vm *VendingMachine) AddProduct(row, column uint, product *product.Product) error {
	return vm.Inventory.AddProduct(row, column, product)
}

func (vm *VendingMachine) RemoveProduct(row, column uint) error {
	return vm.Inventory.RemoveProduct(row, column)
}

func (vm *VendingMachine) GetAllProducts() []*product.Product {
	return vm.Inventory.GetAllProducts()
}

func (vm *VendingMachine) InsertMoney(amount uint) {
	vm.CurrentState.InsertMoney(amount)
}

func (vm *VendingMachine) SelectProduct(row, column uint) {
	vm.CurrentState.SelectProduct(row, column)
}

func (vm *VendingMachine) DispenseProduct(row, column uint) {
	vm.CurrentState.DispenseProduct(row, column)
}

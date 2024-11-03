package main

import (
	"fmt"
	inventory2 "vending-machine/inventory"
	"vending-machine/machine"
	"vending-machine/product"
	"vending-machine/service"
)

func main() {
	inventory := inventory2.NewInventory(5, 5)

	vendingMachine := machine.NewVendingMachine(inventory)

	adminService := service.NewAdminService(vendingMachine)

	err := adminService.AddProduct(0, 0, &product.Product{Name: "Pepsi", Price: 10})
	if err != nil {
		fmt.Println("error adding product ", err)
		return
	}

	err = adminService.AddProduct(0, 1, &product.Product{Name: "Coke", Price: 20})
	if err != nil {
		fmt.Println("error adding product ", err)
		return
	}

	err = adminService.AddProduct(1, 1, &product.Product{Name: "Samosa", Price: 30})
	if err != nil {
		fmt.Println("error adding product ", err)
		return
	}

	err = adminService.AddProduct(1, 0, &product.Product{Name: "Pakoda", Price: 25})
	if err != nil {
		fmt.Println("error adding product ", err)
		return
	}

	adminService.ShowAllProducts()

	userService := service.NewUserService(vendingMachine)
	userService.InsertMoney(40)
	userService.SelectProduct(1, 1)

	adminService.ShowAllProducts()
}

package machine

import (
	"fmt"
)

type DispenseState struct {
	vendingMachine *VendingMachine
}

func NewDispenseState(vendingMachine *VendingMachine) State {
	return &DispenseState{
		vendingMachine: vendingMachine,
	}
}

func (s *DispenseState) InsertMoney(amount uint) {
	fmt.Println("dispensing in progress, please wait")
}

func (s *DispenseState) SelectProduct(row, column uint) {
	fmt.Println("dispensing in progress, please wait")
}

func (s *DispenseState) DispenseProduct(row, column uint) {
	product, err := s.vendingMachine.Inventory.GetProduct(row, column)
	if err != nil {
		fmt.Println("error getting product while dispensing")
	}
	if s.vendingMachine.CurrentAmount >= product.Price {
		change := s.vendingMachine.CurrentAmount - product.Price
		fmt.Printf("dispensing product %s\n", product.Name)
		if change > 0 {
			fmt.Printf("returning change %d\n", change)
		}
		s.vendingMachine.RemoveProduct(row, column)
		s.vendingMachine.CurrentAmount = 0
		s.vendingMachine.CurrentState = NewNoMoneyState(s.vendingMachine)
	} else {
		fmt.Println("dispensing product not available")
	}
}

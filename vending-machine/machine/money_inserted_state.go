package machine

import (
	"fmt"
)

type MoneyInsertedState struct {
	vendingMachine *VendingMachine
}

func NewMoneyInsertedState(vendingMachine *VendingMachine) State {
	return &MoneyInsertedState{
		vendingMachine: vendingMachine,
	}
}

func (s *MoneyInsertedState) InsertMoney(amount uint) {
	fmt.Println("money already inserted, please select product")
}

func (s *MoneyInsertedState) SelectProduct(row, column uint) {
	product, err := s.vendingMachine.Inventory.GetProduct(row, column)
	if err != nil {
		fmt.Println("error getting product - ", err)
		return
	}

	if s.vendingMachine.CurrentAmount < product.Price {
		fmt.Println("money already inserted, please select product")
		s.vendingMachine.CurrentAmount = 0
		s.vendingMachine.CurrentState = NewNoMoneyState(s.vendingMachine)
		return
	}

	fmt.Printf("product selected - %s\n", product.Name)
	s.vendingMachine.CurrentState = NewDispenseState(s.vendingMachine)
	s.vendingMachine.CurrentState.DispenseProduct(row, column)
}

func (s *MoneyInsertedState) DispenseProduct(row, column uint) {
	fmt.Println("please select a product first")
}

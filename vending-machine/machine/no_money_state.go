package machine

import (
	"fmt"
)

type NoMoneyState struct {
	vendingMachine *VendingMachine
}

func NewNoMoneyState(vendingMachine *VendingMachine) State {
	return &NoMoneyState{
		vendingMachine: vendingMachine,
	}
}

func (s *NoMoneyState) InsertMoney(amount uint) {
	fmt.Printf("money inserted: %d\n", amount)
	s.vendingMachine.CurrentAmount = amount
	s.vendingMachine.CurrentState = NewMoneyInsertedState(s.vendingMachine)
}

func (s *NoMoneyState) SelectProduct(row, column uint) {
	fmt.Println("please insert money first")
}

func (s *NoMoneyState) DispenseProduct(row, column uint) {
	fmt.Println("please insert money and select product to dispense")
}

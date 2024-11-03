package service

import "vending-machine/machine"

type UserService struct {
	vendingMachine *machine.VendingMachine
}

func NewUserService(vendingMachine *machine.VendingMachine) *UserService {
	return &UserService{
		vendingMachine: vendingMachine,
	}
}

func (us *UserService) InsertMoney(amount uint) {
	us.vendingMachine.InsertMoney(amount)
}

func (us *UserService) SelectProduct(row, column uint) {
	us.vendingMachine.SelectProduct(row, column)
}

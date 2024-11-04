package game

import (
	"math/rand"
)

type Dice struct{}

func NewDice() *Dice {
	return &Dice{}
}

func (d *Dice) Roll() int {
	return rand.Intn(6) + 1 // returns a number between 1 and 6
}

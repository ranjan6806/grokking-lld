package main

import "snake-and-ladder/game"

func main() {
	g := game.NewGame(100, 2)
	g.Play()
}

package main

import (
	"fmt"
	"tic-tac-toe/board"
	"tic-tac-toe/game"
	"tic-tac-toe/player"
)

func main() {
	var size int
	fmt.Println("Enter board size (n): ")
	fmt.Scan(&size)

	boardObj := board.NewBoard(size)

	var symbol int
	fmt.Print("Player 1, choose your symbol (0 or 1): ")
	fmt.Scan(&symbol)

	player1Obj := player.NewPlayer(symbol)
	player2Obj := player.NewPlayer(1 - symbol)

	gameObj := game.NewGame(boardObj, player1Obj, player2Obj)
	gameObj.Start()
}

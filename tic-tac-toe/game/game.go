package game

import (
	"fmt"
	"tic-tac-toe/board"
	"tic-tac-toe/player"
)

type Game struct {
	board         *board.Board
	player1       *player.Player
	player2       *player.Player
	currentPlayer *player.Player
}

func NewGame(board *board.Board, player1 *player.Player, player2 *player.Player) *Game {
	return &Game{
		board:         board,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
	}
}

func (g *Game) Start() {
	for {
		g.board.Print()
		fmt.Printf("Player %d's turn (enter row and column): ", g.currentPlayer.Symbol)

		var row, col int
		_, err := fmt.Scan(&row, &col)
		if err != nil {
			fmt.Printf("error reading input - %+v", err)
			return
		}

		if g.board.PlaceSymbol(row, col, g.currentPlayer.Symbol) {
			if g.board.CheckWin(g.currentPlayer.Symbol) {
				g.board.Print()
				fmt.Printf("Player %d wins!\n", g.currentPlayer.Symbol)
				return
			}

			if g.board.IsFull() {
				g.board.Print()
				fmt.Println("The game is a draw!")
				return
			}

			if g.currentPlayer == g.player1 {
				g.currentPlayer = g.player2
			} else {
				g.currentPlayer = g.player1
			}

		} else {
			fmt.Println("invalid move, try again")
		}
	}
}

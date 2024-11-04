package game

import (
	"fmt"
	"snake-and-ladder/factory"
)

type Game struct {
	board    *Board
	players  []*Player
	dice     *Dice
	finished bool
}

func NewGame(boardSize, numPlayers int) *Game {
	g := &Game{
		board: NewBoard(boardSize),
		dice:  NewDice(),
	}

	for i := 0; i < numPlayers; i++ {
		player := NewPlayer(fmt.Sprintf("Player%d", i))
		g.players = append(g.players, player)
	}

	g.board.snakes = factory.CreateSnakes()
	g.board.ladders = factory.CreateLadders()
	return g
}

func (g *Game) Play() {
	for !g.finished {
		for _, player := range g.players {
			roll := g.dice.Roll()
			g.board.MovePlayer(player, roll)
			if player.position == g.board.size {
				fmt.Printf("%s wins!\n", player.name)
				g.finished = true
				break
			}
		}
	}
}

package game

import "fmt"

type Board struct {
	size    int
	snakes  map[int]int
	ladders map[int]int
}

func NewBoard(size int) *Board {
	return &Board{
		size:    size,
		snakes:  make(map[int]int),
		ladders: make(map[int]int),
	}
}

func (b *Board) MovePlayer(player *Player, roll int) {
	newPos := player.position + roll
	if newPos > b.size {
		return
	}

	if snakeEnd, ok := b.snakes[newPos]; ok {
		fmt.Printf("%s encountered a snake! moving from %d to %d\n", player.name, newPos, snakeEnd)
		newPos = snakeEnd
	} else if ladderEnd, ok := b.ladders[newPos]; ok {
		fmt.Printf("%s climbed a ladder! moving from %d to %d\n", player.name, newPos, ladderEnd)
	}

	player.position = newPos
}

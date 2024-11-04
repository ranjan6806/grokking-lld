package game

type Player struct {
	name     string
	position int
}

func NewPlayer(name string) *Player {
	return &Player{
		name:     name,
		position: 0,
	}
}

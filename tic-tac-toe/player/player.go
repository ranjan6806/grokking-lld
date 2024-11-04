package player

type Player struct {
	Symbol int
}

func NewPlayer(symbol int) *Player {
	return &Player{Symbol: symbol}
}

package chess

type Piece interface {
	GetType() string
	GetColor() string
	GetPossibleMoves(b *Board, pos Position) []Position
}

type BasePiece struct {
	Type  string
	Color string
}

func (p *BasePiece) GetType() string {
	return p.Type
}

func (p *BasePiece) GetColor() string {
	return p.Color
}

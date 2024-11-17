package chess

type Board struct {
	Grid [8][8]Piece
}

func NewBoard() *Board {
	return &Board{} // pending initialisation
}

func (b *Board) IsValidPosition(pos Position) bool {
	return pos.X >= 0 && pos.X < 8 && pos.Y >= 0 && pos.Y < 8
}

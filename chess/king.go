package chess

type King struct {
	BasePiece
}

func (k *King) GetPossibleMoves(b *Board, pos Position) []Position {
	moves := []Position{}
	directions := []Position{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	for _, direction := range directions {
		newPos := Position{pos.X + direction.X, pos.Y + direction.Y}
		if b.IsValidPosition(newPos) && (b.Grid[newPos.X][newPos.Y] == nil || b.Grid[newPos.X][newPos.Y].GetColor() != k.GetColor()) {
			moves = append(moves, newPos)
		}
	}

	return moves
}

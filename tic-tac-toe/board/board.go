package board

import "fmt"

type Board struct {
	size int
	grid [][]int
}

func NewBoard(size int) *Board {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
		for j := range grid[i] {
			grid[i][j] = -1
		}
	}
	return &Board{size: size, grid: grid}
}

func (b *Board) Print() {
	for _, row := range b.grid {
		for _, cell := range row {
			if cell == -1 {
				fmt.Print(". ")
			} else {
				fmt.Print(cell, " ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) PlaceSymbol(row, column, symbol int) bool {
	if row >= 0 && row < b.size && column >= 0 && column < b.size && b.grid[row][column] == -1 {
		b.grid[row][column] = symbol
		return true
	}
	return false
}

func (b *Board) CheckWin(symbol int) bool {
	for i := 0; i < int(b.size); i++ {
		if b.checkRow(b.grid[i], symbol) || b.checkCol(b.grid, i, symbol) {
			return true
		}
	}

	return b.checkDiagonal(b.grid, symbol)
}

func (b *Board) checkRow(row []int, symbol int) bool {
	for _, cell := range row {
		if cell != symbol {
			return false
		}
	}

	return true
}

func (b *Board) checkCol(grid [][]int, column, symbol int) bool {
	for i := 0; i < len(grid); i++ {
		if grid[i][column] != symbol {
			return false
		}
	}

	return true
}

func (b *Board) checkDiagonal(grid [][]int, symbol int) bool {
	primaryDiagonal, secondaryDiagonal := true, true
	for i := 0; i < len(grid); i++ {
		if grid[i][i] != symbol {
			primaryDiagonal = false
		}

		if grid[i][len(grid)-1-i] != symbol {
			secondaryDiagonal = false
		}
	}

	return primaryDiagonal || secondaryDiagonal
}

func (b *Board) IsFull() bool {
	for _, row := range b.grid {
		for _, cell := range row {
			if cell == -1 {
				return false
			}
		}
	}

	return true
}

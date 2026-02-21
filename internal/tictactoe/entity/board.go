package entity

import (
	"fmt"
	"strings"
)

type board struct {
	size int
	grid [][]string
}

type Board interface {
	GetBoardSize() int
	PlaceMark(row, col int, mark string) error
	String() string
}

func NewBoard(size int) *board {
	grid := make([][]string, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
	}
	return &board{
		size: size,
		grid: grid,
	}
}

func (b *board) GetBoardSize() int {
	return b.size
}

func (b *board) PlaceMark(row, col int, mark string) error {
	if row < 0 || row >= b.size || col < 0 || col >= b.size {
		return fmt.Errorf("invalid move: (%d,%d) out of bounds", row, col)
	}
	if b.grid[row][col] != "" {
		return fmt.Errorf("cell (%d,%d) is already occupied", row, col)
	}
	b.grid[row][col] = mark
	return nil
}

func (b *board) String() string {
	var sb strings.Builder
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == "" {
				sb.WriteString("_ ")
			} else {
				sb.WriteString(b.grid[i][j] + " ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

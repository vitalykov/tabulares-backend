package model

import (
	"errors"
)

type Board[T FigureType] struct {
	Cells [][]Figure[T]
}

func NewBoard[T FigureType](height, width int) (Board[T], error) {
	if height <= 0 || width <= 0 {
		return Board[T]{}, errors.New("Boards dimensions should be positive integers")
	}
	cells := make([][]Figure[T], height)
	for i := range cells {
		cells[i] = make([]Figure[T], width)
	}
	return Board[T]{Cells: cells}, nil
}

func (b Board[T]) Width() int {
	return len(b.Cells[0])
}

func (b Board[T]) Height() int {
	return len(b.Cells)
}

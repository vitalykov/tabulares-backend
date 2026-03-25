package process

import (
	"board-games/internal/domain/model"
)

func getFigure[T model.FigureType](board model.Board[T], row, col int) model.Figure[T] {
	return board.Cells[row][col]
}

func addFigure[T model.FigureType](figure model.Figure[T], board model.Board[T], row, col int) {
	board.Cells[row][col] = figure
}

// Returns the figure which is removed from the board
func removeFigure[T model.FigureType](board model.Board[T], row, col int) model.Figure[T] {
	figure := getFigure(board, row, col)
	board.Cells[row][col] = model.Figure[T]{}
	return figure
}

func makeAction[T model.FigureType](action *model.Action[T], board model.Board[T]) {
	switch action.Type {
	case model.ActionGet:
		action.Figure = getFigure(board, action.Args[0], action.Args[1])
	case model.ActionAdd:
		addFigure(action.Figure, board, action.Args[0], action.Args[1])
	case model.ActionRemove:
		action.Figure = removeFigure(board, action.Args[0], action.Args[1])
	}
}

func undoAction[T model.FigureType](action model.Action[T], board model.Board[T]) {
	switch action.Type {
	case model.ActionAdd:
		_ = removeFigure(board, action.Args[0], action.Args[1])
	case model.ActionRemove:
		addFigure(action.Figure, board, action.Args[0], action.Args[1])
	}
}

package process

import (
	"board-games/internal/domain/model"
	"errors"
)

func MakeMove[T model.FigureType](move model.Move[T], game *model.Game[T]) {
	for i := range move.Actions {
		makeAction(&move.Actions[i], game.Board)
	}
	game.Moves = append(game.Moves, move)
}

func UndoMove[T model.FigureType](game *model.Game[T]) (model.Move[T], error) {
	if len(game.Moves) == 0 {
		return model.Move[T]{}, errors.New("No moves yet. Nothing to undo")
	}
	move := game.Moves[len(game.Moves)-1]
	for i := len(move.Actions) - 1; i >= 0; i-- {
		action := move.Actions[i]
		undoAction(action, game.Board)
	}
	game.Moves = game.Moves[:len(game.Moves)-1]
	game.Turn = move.Player
	game.Winner = model.NoWinner
	return move, nil
}

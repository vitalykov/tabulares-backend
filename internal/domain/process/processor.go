package process

import (
	"board-games/internal/domain/model"
)

type GameProcessor[T model.FigureType] interface {
	GenerateMove(game model.Game[T]) model.Move[T]
	ValidateMove(game model.Game[T], move model.Move[T]) error
	GetWinner(game model.Game[T]) int
	NextTurn(game model.Game[T]) int
	FirstTurn(game model.Game[T]) int
}

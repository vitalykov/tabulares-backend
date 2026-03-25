package service

import (
	"errors"
	"fmt"

	dModel "board-games/internal/domain/model"
	"board-games/internal/domain/process"
	"board-games/internal/usecases/boundaries"
	"board-games/internal/usecases/mappers"
	uModel "board-games/internal/usecases/model"
)

var (
	ErrGameNotFound = errors.New("Game not found")
)

type DefaultGameInteractor[T dModel.FigureType] struct {
	repository boundaries.GameRepository
	processor  process.GameProcessor[T]
}

func NewDefaultGameInteractor[T dModel.FigureType](repo boundaries.GameRepository, processor process.GameProcessor[T]) *DefaultGameInteractor[T] {
	return &DefaultGameInteractor[T]{
		repository: repo,
		processor:  processor,
	}
}

func (gi *DefaultGameInteractor[T]) StartGame(gameInfo *uModel.GameInfo) error {
	if len(gameInfo.Players) == 0 {
		return errors.New("Not enough players")
	}
	if gameInfo.Status != uModel.ReadyToStart {
		msg := fmt.Sprint("Game is: ", gameInfo.Status.String())
		return errors.New(msg)
	}
	game := gameInfo.Game.(*dModel.Game[T])
	if len(game.Moves) == 0 {
		game.Turn = gi.processor.FirstTurn(*game)
	}
	gameInfo.Turn = gameInfo.Players[game.Turn]
	gameInfo.Status = uModel.InProgress
	return nil
}

func (gi *DefaultGameInteractor[T]) MakeMove(gameInfo *uModel.GameInfo, moveInfo uModel.MoveInfo) error {
	if gameInfo.Status == uModel.Finished {
		return errors.New("Game is finished")
	}
	game := gameInfo.Game.(*dModel.Game[T])
	move, err := mappers.ToMove[T](moveInfo, gameInfo)
	if err != nil {
		return err
	}
	if err = gi.processor.ValidateMove(*game, move); err != nil {
		return err
	}

	process.MakeMove(move, game)
	gameInfo.Moves = append(gameInfo.Moves, moveInfo)
	game.Winner = gi.processor.GetWinner(*game)
	if game.Winner != dModel.NoWinner {
		if game.Winner != dModel.Draw {
			gameInfo.Winner = gameInfo.Players[game.Winner]
		}
		gameInfo.Status = uModel.Finished
	}
	game.Turn = gi.processor.NextTurn(*game)
	gameInfo.Turn = gameInfo.Players[game.Turn]
	return nil
}

func (gi *DefaultGameInteractor[T]) UndoMove(gameInfo *uModel.GameInfo) (uModel.MoveInfo, error) {
	game := gameInfo.Game.(*dModel.Game[T])
	if len(gameInfo.Moves) == 0 {
		return uModel.MoveInfo{}, errors.New("No moves yet. Nothing to undo")
	}
	_, err := process.UndoMove(game)
	if err != nil {
		return uModel.MoveInfo{}, err
	}
	moveInfo := gameInfo.Moves[len(gameInfo.Moves)-1]
	gameInfo.Moves = gameInfo.Moves[:len(gameInfo.Moves)-1]
	gameInfo.Turn = gameInfo.Players[game.Turn]
	gameInfo.Winner = uModel.NoWinner
	gameInfo.Status = uModel.InProgress
	return moveInfo, nil
}

func (gi *DefaultGameInteractor[T]) StopGame(gameInfo *uModel.GameInfo) error {
	if len(gameInfo.Moves) == 0 {
		return errors.New("No moves yet. Nothing to stop")
	}
	err := gi.repository.Store(gameInfo)
	if err != nil {
		return err
	}
	gameInfo.Status = uModel.Stopped
	return nil
}

func (gi *DefaultGameInteractor[T]) CancelGame(gameInfo *uModel.GameInfo) error {
	gi.repository.Delete(gameInfo.ID)
	return nil
}

func (gi *DefaultGameInteractor[T]) GetHint(gameInfo *uModel.GameInfo) (uModel.MoveInfo, error) {
	game := gameInfo.Game.(*dModel.Game[T])
	move := gi.processor.GenerateMove(*game)
	return mappers.ToMoveInfo(move, gameInfo), nil
}

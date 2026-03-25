package service

import (
	dModel "board-games/internal/domain/model"
	dService "board-games/internal/domain/process"
	"board-games/internal/usecases/boundaries"
	uModel "board-games/internal/usecases/model"
	"errors"
)

var ErrUnknownGameType = errors.New("Unknown game type")

type GameInteractorSwitch struct {
	repository boundaries.GameRepository
	processors map[uModel.GameType]any
}

func NewGameInteractorSwitch(repo boundaries.GameRepository) *GameInteractorSwitch {
	return &GameInteractorSwitch{
		repository: repo,
		processors: map[uModel.GameType]any{
			uModel.TicTacToeType: NewDefaultGameInteractor(repo, dService.TicTacToeProcessor{}),
		},
	}
}

func (gs *GameInteractorSwitch) StartGame(gameInfo *uModel.GameInfo) error {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		return gs.processors[gameInfo.Type].(*DefaultGameInteractor[dModel.TicTacToeType]).StartGame(gameInfo)
	}
	return ErrUnknownGameType
}

func (gs *GameInteractorSwitch) StopGame(gameInfo *uModel.GameInfo) error {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		return gs.processors[gameInfo.Type].(*DefaultGameInteractor[dModel.TicTacToeType]).StopGame(gameInfo)
	}
	return ErrUnknownGameType
}

func (gs *GameInteractorSwitch) CancelGame(gameInfo *uModel.GameInfo) error {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		return gs.processors[gameInfo.Type].(*DefaultGameInteractor[dModel.TicTacToeType]).CancelGame(gameInfo)
	}
	return ErrUnknownGameType
}

func (gs *GameInteractorSwitch) MakeMove(gameInfo *uModel.GameInfo, moveInfo uModel.MoveInfo) error {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		return gs.processors[gameInfo.Type].(*DefaultGameInteractor[dModel.TicTacToeType]).MakeMove(gameInfo, moveInfo)
	}
	return ErrUnknownGameType
}

func (gs *GameInteractorSwitch) UndoMove(gameInfo *uModel.GameInfo) (uModel.MoveInfo, error) {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		return gs.processors[gameInfo.Type].(*DefaultGameInteractor[dModel.TicTacToeType]).UndoMove(gameInfo)
	}
	return uModel.MoveInfo{}, ErrUnknownGameType
}

func (gs *GameInteractorSwitch) GetHint(gameInfo *uModel.GameInfo) (uModel.MoveInfo, error) {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		return gs.processors[gameInfo.Type].(*DefaultGameInteractor[dModel.TicTacToeType]).GetHint(gameInfo)
	}
	return uModel.MoveInfo{}, ErrUnknownGameType
}

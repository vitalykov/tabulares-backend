package service

import (
	dModel "board-games/internal/domain/model"
	"board-games/internal/usecases/boundaries"
	"board-games/internal/usecases/mappers"
	uModel "board-games/internal/usecases/model"
	"errors"
	"slices"

	"github.com/google/uuid"
)

type GameMaster struct {
	repository boundaries.GameRepository
}

func NewGameMaster(repository boundaries.GameRepository) *GameMaster {
	return &GameMaster{
		repository: repository,
	}
}

func (gm GameMaster) CreateGame(newGameInfo uModel.NewGameInfo) (*uModel.GameInfo, error) {
	gameUUID := uuid.New()
	var game any
	switch newGameInfo.Type {
	case uModel.TicTacToeType:
		game = mappers.GetGame[dModel.TicTacToeType](newGameInfo)
	}
	gameInfo := &uModel.GameInfo{
		ID:             gameUUID,
		Type:           newGameInfo.Type,
		BoardWidth:     newGameInfo.BoardWidth,
		BoardHeight:    newGameInfo.BoardHeight,
		Players:        newGameInfo.Players,
		Moves:          make([]uModel.MoveInfo, 0),
		Winner:         uModel.NoWinner,
		Status:         uModel.ReadyToStart,
		AdditionalInfo: newGameInfo.AdditionalInfo,
		Game:           game,
	}
	if err := gm.repository.Store(gameInfo); err != nil {
		return nil, err
	}
	return gameInfo, nil
}

func (gm GameMaster) LoadGame(gameUUID uModel.UUID) (*uModel.GameInfo, error) {
	gameInfo, err := gm.repository.Load(gameUUID)
	if err != nil {
		return nil, err
	}
	return gameInfo, nil
}

func (gm GameMaster) AddPlayer(playerID uModel.PlayerID, gameInfo *uModel.GameInfo) error {
	if slices.Contains(gameInfo.Players, playerID) {
		return errors.New("Player already in the game")
	}
	gameInfo.Players = append(gameInfo.Players, playerID)
	return nil
}

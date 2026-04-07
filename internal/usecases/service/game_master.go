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
	cache boundaries.GameCacheRepository
	db    boundaries.GameRepository
}

func NewGameMaster(cache boundaries.GameCacheRepository, db boundaries.GameRepository) *GameMaster {
	return &GameMaster{
		cache: cache,
		db:    db,
	}
}

func (gm GameMaster) CreateGame(newGameInfo uModel.NewGameInfo) (*uModel.GameInfo, error) {
	gameUUID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
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
	if err := gm.cache.Store(gameInfo); err != nil {
		return nil, err
	}
	if err := gm.db.Store(gameInfo); err != nil {
		return nil, err
	}
	return gameInfo, nil
}

func (gm GameMaster) LoadGame(gameUUID uModel.UUID) (*uModel.GameInfo, error) {
	gameInfo, err := gm.cache.Load(gameUUID)
	if err == nil {
		return gameInfo, nil
	}
	gameInfo, err = gm.db.Load(gameUUID)
	if err != nil {
		return nil, err
	}
	if err := gm.cache.Store(gameInfo); err != nil {
		return nil, err
	}
	return gameInfo, nil
}

// Temporary unused
func (gm GameMaster) AddPlayer(playerID uModel.PlayerID, gameInfo *uModel.GameInfo) error {
	if slices.Contains(gameInfo.Players, playerID) {
		return errors.New("Player already in the game")
	}
	gameInfo.Players = append(gameInfo.Players, playerID)
	return nil
}

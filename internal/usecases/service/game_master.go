package service

import (
	dModel "board-games/internal/domain/model"
	"board-games/internal/usecases/boundaries"
	"board-games/internal/usecases/mappers"
	uModel "board-games/internal/usecases/model"
	"errors"

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
		Winner:         uModel.NoWinnerID,
		Status:         uModel.GameWaitingForPlayers,
		AdditionalInfo: newGameInfo.AdditionalInfo,
		Game:           game,
	}
	if err := gm.cache.StoreGame(gameInfo); err != nil {
		return nil, err
	}
	return gameInfo, nil
}

func (gm GameMaster) LoadGame(gameUUID uModel.GameID) (*uModel.GameInfo, error) {
	gameInfo, err := gm.cache.GetGame(gameUUID)
	if err == nil {
		return gameInfo, nil
	}
	gameInfo, err = gm.db.Load(gameUUID)
	if err != nil {
		return nil, err
	}
	if err := gm.cache.StoreGame(gameInfo); err != nil {
		return nil, err
	}
	return gameInfo, nil
}

const noFreeSlot = -1

func findPlayer(playerID uModel.PlayerID, players []uModel.PlayerInfo) (bool, int) {
	for i, playerInfo := range players {
		if playerInfo.PlayerID == playerID {
			return true, i
		}
		if playerInfo.PlayerID == uModel.NoWinnerID {
			return false, i
		}
	}
	return false, noFreeSlot
}

func (gm GameMaster) AddPlayer(playerID uModel.PlayerID, gameInfo *uModel.GameInfo) error {
	if gameInfo.Status != uModel.GameWaitingForPlayers {
		return errors.New("Game is not waiting for players")
	}
	ok, i := findPlayer(playerID, gameInfo.Players)
	if i == noFreeSlot {
		return errors.New("No free slot for player")
	}
	if ok {
		gameInfo.Players[i].Status = uModel.PlayerReady
	} else {
		gameInfo.Players[i] = uModel.PlayerInfo{
			PlayerID: playerID,
			Status:   uModel.PlayerReady,
		}
	}
	return nil
}

func (gm GameMaster) GetOngoingGames() ([]*uModel.GameInfo, error) {
	return gm.cache.GetAllGames()
}

func (gm GameMaster) GetPlayerGames(playerID uModel.PlayerID) ([]*uModel.GameInfo, error) {
	return gm.cache.GetPlayerGames(playerID)
}

package cache

import (
	"board-games/internal/usecases/model"
	"board-games/pkg/cache"
	"errors"
)

const cacheCapacity = 1_000_000

type GameCache struct {
	games   *cache.LRUCache[model.GameID, *model.GameInfo]
	players *cache.LRUCache[model.PlayerID, []*model.GameInfo]
}

func NewGameCache() (*GameCache, error) {
	games, err := cache.NewLRUCache[model.GameID, *model.GameInfo](cacheCapacity)
	if err != nil {
		return nil, err
	}
	players, err := cache.NewLRUCache[model.PlayerID, []*model.GameInfo](cacheCapacity)
	if err != nil {
		return nil, err
	}
	return &GameCache{
		games:   games,
		players: players,
	}, nil
}

func (gc *GameCache) GetGame(gameID model.GameID) (*model.GameInfo, error) {
	gameInfo, ok := gc.games.Get(gameID)
	if !ok {
		return nil, errors.New("Not found in cache")
	}
	return gameInfo, nil
}

func (gc *GameCache) GetPlayerGames(playerID model.PlayerID) ([]*model.GameInfo, error) {
	gamesList, ok := gc.players.Get(playerID)
	if !ok {
		return nil, errors.New("Not found in cache")
	}
	return gamesList, nil
}

func (gc *GameCache) GetAllGames() ([]*model.GameInfo, error) {
	return gc.games.GetAll(), nil
}

func (gc *GameCache) StoreGame(gameInfo *model.GameInfo) error {
	gc.games.Set(gameInfo.ID, gameInfo)
	for _, playerInfo := range gameInfo.Players {
		gamesList, ok := gc.players.Get(playerInfo.PlayerID)
		if !ok {
			gamesList = make([]*model.GameInfo, 0)
		}
		gamesList = append(gamesList, gameInfo)
		gc.players.Set(playerInfo.PlayerID, gamesList)
	}
	return nil
}

func (gc *GameCache) DeleteGame(gameID model.GameID) error {
	gameInfo, ok := gc.games.Get(gameID)
	if !ok {
		return errors.New("Not found in cache")
	}
	for _, playerInfo := range gameInfo.Players {
		gamesList, ok := gc.players.Get(playerInfo.PlayerID)
		if !ok {
			return errors.New("Not found in cache")
		}
		for i, game := range gamesList {
			if game.ID == gameID {
				gamesList = append(gamesList[:i], gamesList[i+1:]...)
				break
			}
		}
		gc.players.Set(playerInfo.PlayerID, gamesList)
	}
	gc.games.Delete(gameID)
	return nil
}

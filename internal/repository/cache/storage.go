package cache

import (
	"board-games/internal/usecases/model"
	"board-games/pkg/cache"
	"errors"
)

const cacheCapacity = 1_000_000

type GameCache struct {
	cache *cache.LRUCache[model.UUID, *model.GameInfo]
}

func NewGameCache() (*GameCache, error) {
	cache, err := cache.NewLRUCache[model.UUID, *model.GameInfo](cacheCapacity)
	if err != nil {
		return nil, err
	}
	return &GameCache{
		cache: cache,
	}, nil
}

func (gc *GameCache) Load(gameID model.UUID) (*model.GameInfo, error) {
	gameInfo, ok := gc.cache.Get(gameID)
	if !ok {
		return nil, errors.New("Not found in cache")
	}
	return gameInfo, nil
}

func (gc *GameCache) Store(gameInfo *model.GameInfo) error {
	gc.cache.Set(gameInfo.ID, gameInfo)
	return nil
}

func (gc *GameCache) Delete(gameID model.UUID) error {
	gc.cache.Delete(gameID)
	return nil
}

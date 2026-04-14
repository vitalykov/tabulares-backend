package boundaries

import "board-games/internal/usecases/model"

type GameCacheRepository interface {
	StoreGame(gameInfo *model.GameInfo) error
	GetGame(gameID model.UUID) (*model.GameInfo, error)
	GetAllGames() ([]*model.GameInfo, error)
	GetPlayerGames(playerID model.PlayerID) ([]*model.GameInfo, error)
	DeleteGame(gameID model.UUID) error
}

type GameRepository interface {
	Store(gameInfo *model.GameInfo) error
	Load(gameID model.UUID) (*model.GameInfo, error)
	FindByPlayerID(playerID model.PlayerID) ([]*model.GameInfo, error)
	Delete(gameID model.UUID) error
}

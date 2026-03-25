package boundaries

import "board-games/internal/usecases/model"

type GameRepository interface {
	Store(gameInfo *model.GameInfo) error
	Load(gameID model.UUID) (*model.GameInfo, error)
	Delete(gameID model.UUID) error
}

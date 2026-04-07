package pg

import (
	"board-games/internal/usecases/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGGameRepository struct {
	db *pgxpool.Pool
}

func NewPGGameRepository(db *pgxpool.Pool) *PGGameRepository {
	return &PGGameRepository{db: db}
}

func (r *PGGameRepository) Store(gameInfo *model.GameInfo) error {
	return nil
}

func (r *PGGameRepository) Load(gameID model.UUID) (*model.GameInfo, error) {
	return nil, nil
}

func (r *PGGameRepository) FindByPlayerID(playerID model.PlayerID) ([]*model.GameInfo, error) {
	return nil, nil
}

func (r *PGGameRepository) Delete(gameID model.UUID) error {
	return nil
}

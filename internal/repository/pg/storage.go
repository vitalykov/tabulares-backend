package pg

import (
	"board-games/internal/usecases/model"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGGameRepository struct {
	db *pgxpool.Pool
}

func NewPGGameRepository(db *pgxpool.Pool) *PGGameRepository {
	return &PGGameRepository{db: db}
}

func (r *PGGameRepository) Store(gameInfo *model.GameInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := r.db.Exec(ctx,
		`INSERT INTO games (game_id, type, board_width, board_height, winner, additional_info)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		gameInfo.ID, gameInfo.Type, gameInfo.BoardWidth, gameInfo.BoardHeight, gameInfo.Winner, gameInfo.AdditionalInfo,
	)
	if err != nil {
		return fmt.Errorf("PGGameRepository.Store: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("PGGameRepository.Store.RowsAffected: %w", err)
	}

	for i, playerID := range gameInfo.Players {
		result, err = r.db.Exec(ctx,
			`INSERT INTO players (game_id, player_id, position)
			 VALUES ($1, $2, $3)`,
			gameInfo.ID, playerID, i,
		)
		if err != nil {
			return fmt.Errorf("PGGameRepository.Store: %w", err)
		}
		if result.RowsAffected() == 0 {
			return fmt.Errorf("PGGameRepository.Store.RowsAffected: %w", err)
		}
	}
	for i, moveInfo := range gameInfo.Moves {
		result, err = r.db.Exec(ctx,
			`INSERT INTO moves (game_id, player_id, move, position)
			 VALUES ($1, $2, $3, $4)`,
			gameInfo.ID, moveInfo.PlayerID, moveInfo.MoveRepr, i,
		)
		if err != nil {
			return fmt.Errorf("PGGameRepository.Store: %w", err)
		}
		if result.RowsAffected() == 0 {
			return fmt.Errorf("PGGameRepository.Store.RowsAffected: %w", err)
		}
	}
	return nil
}

func (r *PGGameRepository) Load(gameID model.GameID) (*model.GameInfo, error) {
	return nil, nil
}

func (r *PGGameRepository) FindByPlayerID(playerID model.PlayerID) ([]*model.GameInfo, error) {
	return nil, nil
}

func (r *PGGameRepository) Delete(gameID model.GameID) error {
	return nil
}

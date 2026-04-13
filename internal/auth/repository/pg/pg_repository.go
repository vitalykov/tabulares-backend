package pg

import (
	"board-games/internal/auth/repository/mappers"
	rModel "board-games/internal/auth/repository/model"
	uModel "board-games/internal/auth/usecase/model"
	dModel "board-games/internal/domain/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgUserRepo struct {
	pool *pgxpool.Pool
}

func NewPgUserRepo(pool *pgxpool.Pool) *PgUserRepo {
	return &PgUserRepo{
		pool: pool,
	}
}

func (r *PgUserRepo) Register(user *uModel.User) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("PgUserRepo.Register: begin: %w", err)
	}
	defer tx.Rollback(context.Background())
	if err := tx.QueryRow(context.Background(), `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING user_id`,
		user.Username,
		user.Password,
	).Scan(&user.UserID); err != nil {
		return fmt.Errorf("PgUserRepo.Register: operation: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("PgUserRepo.Register: commit: %w", err)
	}
	return nil
}

func (r *PgUserRepo) Update(user *uModel.User) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("PgUserRepo.Update: begin: %w", err)
	}
	defer tx.Rollback(context.Background())
	if _, err := tx.Exec(context.Background(), `
		UPDATE users
		SET username = $1, password = $2
		WHERE user_id = $4`,
		user.Username,
		user.Password,
		user.UserID,
	); err != nil {
		return fmt.Errorf("PgUserRepo.Update: operation: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("PgUserRepo.Update: commit: %w", err)
	}
	return nil
}

func (r *PgUserRepo) GetByID(userID dModel.UserID) (*uModel.User, error) {
	var userData rModel.UserData
	err := r.pool.QueryRow(context.Background(), `
		SELECT user_id, username, password, created_at
		FROM users
		WHERE user_id = $1`,
		userID,
	).Scan(
		&userData.UserID,
		&userData.Username,
		&userData.Password,
		&userData.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("PgUserRepo.GetByID: %w", err)
	}
	return mappers.ToUser(&userData), nil
}

func (r *PgUserRepo) FindByUsername(username string) (*uModel.User, error) {
	var userData rModel.UserData
	err := r.pool.QueryRow(context.Background(), `
		SELECT user_id, username, password, created_at
		FROM users
		WHERE username = $1`,
		username,
	).Scan(
		&userData.UserID,
		&userData.Username,
		&userData.Password,
		&userData.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("PgUserRepo.FindByUsername: %w", err)
	}
	return mappers.ToUser(&userData), nil
}

func (r *PgUserRepo) Delete(userID dModel.UserID) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("PgUserRepo.Delete: begin: %w", err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `
		DELETE FROM users
		WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return fmt.Errorf("PgUserRepo.Delete: operation: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("PgUserRepo.Delete: commit: %w", err)
	}
	return nil
}

package repository

import (
	"board-games/internal/domain/model"
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

func (r *PgUserRepo) Register(user *model.User) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("PgUserRepo.Register: %w", err)
	}
	defer tx.Rollback(context.Background())
	if err := tx.QueryRow(context.Background(), `
		INSERT INTO users (username, password, created_at)
		VALUES ($1, $2, $3)
		RETURNING user_id`,
		user.Username,
		user.Password,
		user.CreatedAt,
	).Scan(&user); err != nil {
		return fmt.Errorf("PgUserRepo.Register: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("PgUserRepo.Register: %w", err)
	}
	return nil
}

func (r *PgUserRepo) Update(user *model.User) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("PgUserRepo.Update: %w", err)
	}
	defer tx.Rollback(context.Background())
	if _, err := tx.Exec(context.Background(), `
		UPDATE users
		SET username = $1, password = $2, created_at = $3
		WHERE user_id = $4`,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UserID,
	); err != nil {
		return fmt.Errorf("PgUserRepo.Update: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("PgUserRepo.Update: %w", err)
	}
	return nil
}

func (r *PgUserRepo) GetByID(userID model.UserID) (*model.User, error) {
	var user model.User
	err := r.pool.QueryRow(context.Background(), `
		SELECT user_id, username, password, created_at
		FROM users
		WHERE user_id = $1`,
		userID,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("PgUserRepo.GetByID: %w", err)
	}
	return &user, nil
}

func (r *PgUserRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.pool.QueryRow(context.Background(), `
		SELECT user_id, username, password, created_at
		FROM users
		WHERE username = $1`,
		username,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("PgUserRepo.FindByUsername: %w", err)
	}
	return &user, nil
}

func (r *PgUserRepo) Delete(userID model.UserID) error {
	_, err := r.pool.Exec(context.Background(), `
		DELETE FROM users
		WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return fmt.Errorf("PgUserRepo.Delete: %w", err)
	}
	return nil
}

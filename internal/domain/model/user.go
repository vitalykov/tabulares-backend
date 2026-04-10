package model

import (
	"time"

	"board-games/pkg/pass"

	"github.com/google/uuid"
)

type (
	UserID = uuid.UUID
	GameID = uuid.UUID
)

type User struct {
	UserID    UserID    `json:"user_id"`
	Username  string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) HashPassword(password string) error {
	hash, err := pass.Argon2Hash(password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) VerifyPassword(password string) (bool, error) {
	return pass.Argon2Verify(password, u.Password)
}

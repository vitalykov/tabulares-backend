package model

import (
	"time"

	"board-games/pkg/pass"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Login     string    `json:"login"`
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

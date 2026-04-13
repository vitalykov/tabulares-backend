package model

import (
	"board-games/pkg/pass"

	"board-games/internal/domain/model"
)

type User struct {
	UserID   model.UserID
	Username string
	Password string
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

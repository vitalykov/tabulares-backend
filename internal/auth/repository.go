package auth

import (
	uModel "board-games/internal/auth/usecase/model"
	dModel "board-games/internal/domain/model"
)

type UserRepo interface {
	Register(user *uModel.User) error
	Update(user *uModel.User) error
	GetByID(userID dModel.UserID) (*uModel.User, error)
	FindByUsername(username string) (*uModel.User, error)
	Delete(userID dModel.UserID) error
}

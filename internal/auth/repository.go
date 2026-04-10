package auth

import "board-games/internal/domain/model"

type UserRepo interface {
	Register(user *model.User) error
	Update(user *model.User) error
	GetByID(userID model.UserID) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Delete(userID model.UserID) error
}

package usecase

import (
	"board-games/internal/auth"
	"board-games/internal/domain/model"
)

type UserManager struct {
	repo auth.UserRepo
}

func NewUserManager(repo auth.UserRepo) *UserManager {
	return &UserManager{
		repo: repo,
	}
}

func (um *UserManager) Register(user *model.User) error {
	return um.repo.Register(user)
}

func (um *UserManager) Update(user *model.User) error {
	return um.repo.Update(user)
}

func (um *UserManager) GetByID(userID model.UserID) (*model.User, error) {
	return um.repo.GetByID(userID)
}

func (um *UserManager) FindByUsername(username string) (*model.User, error) {
	return um.repo.FindByUsername(username)
}

func (um *UserManager) Delete(userID model.UserID) error {
	return um.repo.Delete(userID)
}

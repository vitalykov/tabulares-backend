package service

import (
	"board-games/internal/auth"
	uModel "board-games/internal/auth/usecase/model"
	dModel "board-games/internal/domain/model"
	"fmt"
)

type UserManager struct {
	repo auth.UserRepo
}

func NewUserManager(repo auth.UserRepo) *UserManager {
	return &UserManager{
		repo: repo,
	}
}

func (um *UserManager) Register(user *uModel.User) error {
	if err := user.HashPassword(user.Password); err != nil {
		return fmt.Errorf("UserManager.Register: %w", err)
	}
	return um.repo.Register(user)
}

func (um *UserManager) Login(user *uModel.User) error {
	foundUser, err := um.repo.FindByUsername(user.Username)
	if err != nil {
		return fmt.Errorf("UserManager.Login: %w", err)
	}
	ok, err := foundUser.VerifyPassword(user.Password)
	if err != nil {
		return fmt.Errorf("UserManager.Login: %w", err)
	}
	if !ok {
		return fmt.Errorf("UserManager.Login: wrong password")
	}
	user.UserID = foundUser.UserID
	return nil
}

func (um *UserManager) Update(user *uModel.User) error {
	return um.repo.Update(user)
}

func (um *UserManager) GetByID(userID dModel.UserID) (*uModel.User, error) {
	return um.repo.GetByID(userID)
}

func (um *UserManager) FindByUsername(username string) (*uModel.User, error) {
	return um.repo.FindByUsername(username)
}

func (um *UserManager) Delete(userID dModel.UserID) error {
	return um.repo.Delete(userID)
}

package mappers

import (
	rModel "board-games/internal/auth/repository/model"
	uModel "board-games/internal/auth/usecase/model"
)

func ToUser(user *rModel.UserData) *uModel.User {
	return &uModel.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
}

func ToUserData(user *uModel.User) *rModel.UserData {
	return &rModel.UserData{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
}

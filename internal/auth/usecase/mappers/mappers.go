package mappers

import (
	uModel "board-games/internal/auth/usecase/model"
	dModel "board-games/internal/domain/model"
)

func ToUserID(user *uModel.User) dModel.UserID {
	return user.UserID
}

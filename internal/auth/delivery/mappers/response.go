package mappers

import (
	"encoding/json"

	dModel "board-games/internal/auth/delivery/model"
	uModel "board-games/internal/auth/usecase/model"
)

func ToRegisterResponse(userID string) ([]byte, error) {
	return json.Marshal(dModel.RegisterResponse{
		UserID: userID,
	})
}

func ToLoginResponse(userID string) ([]byte, error) {
	return json.Marshal(dModel.LoginResponse{
		UserID: userID,
	})
}

func ToGetByIDResponse(user *uModel.User) ([]byte, error) {
	return json.Marshal(dModel.GetByIDResponse{
		UserID:   user.UserID.String(),
		Username: user.Username,
	})
}

func ToFindByUsernameResponse(user *uModel.User) ([]byte, error) {
	return json.Marshal(dModel.FindByUsernameResponse{
		UserID:   user.UserID.String(),
		Username: user.Username,
	})
}

func ToDeleteResponse(userID string) ([]byte, error) {
	return json.Marshal(dModel.DeleteResponse{
		UserID: userID,
	})
}

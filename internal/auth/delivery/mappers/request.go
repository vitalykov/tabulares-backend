package mappers

import (
	dModel "board-games/internal/auth/delivery/model"
	uModel "board-games/internal/auth/usecase/model"
	"encoding/json"
	"errors"
)

func ToRegisterRequest(data []byte) (*dModel.RegisterRequest, error) {
	var input dModel.RegisterRequest
	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}
	if input.Username == "" {
		return nil, errors.New("username is empty")
	}
	if input.Password == "" {
		return nil, errors.New("password is empty")
	}
	return &input, nil
}

func ToLoginRequest(data []byte) (*dModel.LoginRequest, error) {
	var input dModel.LoginRequest
	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}
	if input.Username == "" {
		return nil, errors.New("username is empty")
	}
	if input.Password == "" {
		return nil, errors.New("password is empty")
	}
	return &input, nil
}

func FromRegisterRequestToUser(input *dModel.RegisterRequest) *uModel.User {
	return &uModel.User{
		Username: input.Username,
		Password: input.Password,
	}
}

func FromLoginRequestToUser(input *dModel.LoginRequest) *uModel.User {
	return &uModel.User{
		Username: input.Username,
		Password: input.Password,
	}
}

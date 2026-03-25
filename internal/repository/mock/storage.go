package mock

import (
	"board-games/internal/usecases/model"
)

type MockRepository struct{}

func (mr MockRepository) Store(_ *model.GameInfo) error {
	return nil
}

func (mr MockRepository) Load(_ model.UUID) (*model.GameInfo, error) {
	return nil, nil
}

func (mr MockRepository) Delete(_ model.UUID) error {
	return nil
}

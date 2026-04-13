package model

import (
	"time"

	"board-games/internal/domain/model"
)

type UserData struct {
	UserID    model.UserID
	Username  string
	Password  string
	CreatedAt time.Time
}

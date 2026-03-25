package model

import "board-games/internal/usecases/model"

type MoveRequest struct {
	PlayerID model.PlayerID `json:"player_id"`
	Move     string         `json:"move"`
}

type NewGameRequest struct {
	Name           string           `json:"name"`
	Players        []model.PlayerID `json:"players"`
	BoardWidth     int              `json:"board_width"`
	BoardHeight    int              `json:"board_height"`
	AdditionalInfo string           `json:"additional_info"`
}

package model

import "board-games/internal/usecases/model"

type GameResponse struct {
	Status   string         `json:"status"`
	Turn     model.PlayerID `json:"turn"`
	Winner   model.PlayerID `json:"winner"`
	PlayerID model.PlayerID `json:"player_id,omitempty"`
	Move     string         `json:"move,omitempty"`
}

type GameResponseFull struct {
	ID             model.UUID       `json:"id"`
	Name           string           `json:"name"`
	BoardWidth     int              `json:"board_width"`
	BoardHeight    int              `json:"board_height"`
	Players        []model.PlayerID `json:"players"`
	Moves          []model.MoveInfo `json:"moves"`
	Winner         model.PlayerID   `json:"winner"`
	Status         string           `json:"status"`
	Turn           model.PlayerID   `json:"turn"`
	AdditionalInfo string           `json:"additional_info,omitempty"`
}

type NewGameResponse struct {
	ID             model.UUID       `json:"id"`
	Name           string           `json:"name"`
	Players        []model.PlayerID `json:"players"`
	BoardWidth     int              `json:"board_width"`
	BoardHeight    int              `json:"board_height"`
	AdditionalInfo string           `json:"additional_info,omitempty"`
}

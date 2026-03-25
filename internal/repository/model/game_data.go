package model

import "board-games/internal/usecases/model"

type GameData struct {
	ID             model.UUID
	NameID         int
	BoardWidth     int
	BoardHeight    int
	Players        []model.PlayerID
	Moves          string
	Winner         model.PlayerID
	AdditionalInfo string
}

package model

type MoveInfo struct {
	PlayerID PlayerID `json:"player_id"`
	MoveRepr string   `json:"move"`
}

type NewGameInfo struct {
	Type           GameType
	BoardWidth     int
	BoardHeight    int
	Players        []PlayerID
	AdditionalInfo string
}

type GameInfo struct {
	ID             UUID
	Type           GameType
	BoardWidth     int
	BoardHeight    int
	Players        []PlayerID
	Moves          []MoveInfo
	Winner         PlayerID
	Status         GameStatus
	Turn           PlayerID
	AdditionalInfo string
	Game           any
}

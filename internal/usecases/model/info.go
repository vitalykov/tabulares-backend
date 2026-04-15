package model

type MoveInfo struct {
	PlayerID PlayerID `json:"player_id"`
	MoveRepr string   `json:"move"`
}

type PlayerInfo struct {
	PlayerID PlayerID     `json:"player_id"`
	Status   PlayerStatus `json:"status"`
}

type NewGameInfo struct {
	Type           GameType
	BoardWidth     int
	BoardHeight    int
	Players        []PlayerInfo
	AdditionalInfo string
}

type GameInfo struct {
	ID             GameID
	Type           GameType
	BoardWidth     int
	BoardHeight    int
	Players        []PlayerInfo
	Moves          []MoveInfo
	Winner         PlayerID
	Status         GameStatus
	Turn           PlayerID
	AdditionalInfo string
	Game           any
}

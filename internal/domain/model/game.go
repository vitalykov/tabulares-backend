package model

type Move[T FigureType] struct {
	Player  int
	Actions []Action[T]
}

type Game[T FigureType] struct {
	NumPlayers     int
	Board          Board[T]
	Moves          []Move[T]
	Turn           int
	Winner         int
	Finished       bool
	AdditionalInfo string
}

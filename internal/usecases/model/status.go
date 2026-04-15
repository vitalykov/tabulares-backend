package model

type GameStatus int

const (
	ReadyToStart GameStatus = iota
	WaitingForPlayers
	InProgress
	Stopped
	Finished
)

func (s GameStatus) Int() int {
	return int(s)
}

func (s GameStatus) String() string {
	return [...]string{
		"Ready to start",
		"Waiting for players",
		"In progress",
		"Stopped",
		"Finished",
	}[s.Int()]
}

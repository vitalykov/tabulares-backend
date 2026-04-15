package model

type GameStatus int

const (
	GameReadyToStart GameStatus = iota
	GameWaitingForPlayers
	GameInProgress
	GameStopped
	GameFinished
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

type PlayerStatus int

const (
	PlayerReady PlayerStatus = iota
	PlayerWaiting
)

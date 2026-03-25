package model

type GameStatus int

const (
	ReadyToStart GameStatus = iota
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
		"In progress",
		"Stopped",
		"Finished",
	}[s.Int()]
}

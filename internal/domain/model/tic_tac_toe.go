package model

type TicTacToeType int

const (
	TicTacToeNoFigure TicTacToeType = iota
	TicTacToeMark
)

func (t TicTacToeType) Int() int {
	return int(t)
}

func (t TicTacToeType) String() string {
	return [...]string{
		" ",
		"Mark",
	}[t.Int()]
}

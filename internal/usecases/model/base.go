package model

import (
	"board-games/internal/domain/model"
	"strconv"

	"github.com/google/uuid"
)

type UUID = uuid.UUID

// PlayerID must be positive, zero and negative values are reserved
type PlayerID int64

func (p PlayerID) String() string {
	return strconv.FormatInt(int64(p), 10)
}

const NoWinner PlayerID = 0

type GameType int

const (
	TicTacToeType GameType = iota
)

func (t GameType) Int() int {
	return int(t)
}

func (t GameType) String() string {
	return [...]string{
		"tic-tac-toe",
	}[t.Int()]
}

func (t GameType) FigureType() model.FigureType {
	return [...]model.FigureType{
		model.TicTacToeNoFigure,
	}[t.Int()]
}

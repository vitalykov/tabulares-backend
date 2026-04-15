package model

import (
	"board-games/internal/domain/model"

	"github.com/google/uuid"
)

type (
	GameID   = uuid.UUID
	PlayerID = uuid.UUID
)

var (
	NoWinnerID PlayerID = uuid.Nil
	DrawID     PlayerID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
)

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

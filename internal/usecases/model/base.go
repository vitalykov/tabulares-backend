package model

import (
	"board-games/internal/domain/model"

	"github.com/google/uuid"
)

type UUID = uuid.UUID

type PlayerID = uuid.UUID

var NoWinner PlayerID = uuid.Nil

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

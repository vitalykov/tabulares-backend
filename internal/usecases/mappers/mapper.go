package mappers

import (
	dModel "board-games/internal/domain/model"
	uModel "board-games/internal/usecases/model"
	"errors"
	"fmt"
)

// var Games = map[uModel.UUID]any{}

func ToGame[T dModel.FigureType](gameInfo *uModel.GameInfo) *dModel.Game[T] {
	return gameInfo.Game.(*dModel.Game[T])
}

func GetGame[T dModel.FigureType](gameInfo uModel.NewGameInfo) *dModel.Game[T] {
	board, _ := dModel.NewBoard[T](gameInfo.BoardHeight, gameInfo.BoardWidth)
	return &dModel.Game[T]{
		NumPlayers:     len(gameInfo.Players),
		Board:          board,
		Moves:          make([]dModel.Move[T], 0),
		Turn:           0,
		Winner:         dModel.NoWinner,
		Finished:       false,
		AdditionalInfo: gameInfo.AdditionalInfo,
	}
}

func ToMove[T dModel.FigureType](moveInfo uModel.MoveInfo, gameInfo *uModel.GameInfo) (dModel.Move[T], error) {
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		move, err := ticTacToeToMove(moveInfo, gameInfo)
		return move.(dModel.Move[T]), err
	}
	return dModel.Move[T]{}, errors.New("Unknown Game Type")
}

func ticTacToeToMove(moveInfo uModel.MoveInfo, gameInfo *uModel.GameInfo) (any, error) {
	var row, col int
	_, err := fmt.Sscan(moveInfo.MoveRepr, &row, &col)
	if err != nil {
		return dModel.Move[dModel.TicTacToeType]{}, err
	}
	player := -1
	for i := 0; i < len(gameInfo.Players); i++ {
		if gameInfo.Players[i] == moveInfo.PlayerID {
			player = i
			break
		}
	}
	if player == -1 {
		return dModel.Move[dModel.TicTacToeType]{}, errors.New("Unknown player")
	}
	return dModel.Move[dModel.TicTacToeType]{
		Player: player,
		Actions: []dModel.Action[dModel.TicTacToeType]{
			{
				Type: dModel.ActionAdd,
				Figure: dModel.Figure[dModel.TicTacToeType]{
					Player: player,
					Type:   dModel.TicTacToeMark,
				},
				Args: []int{row, col},
			},
		},
	}, nil
}

func ToMoveInfo[T dModel.FigureType](move dModel.Move[T], gameInfo *uModel.GameInfo) uModel.MoveInfo {
	var moveRepr string
	switch gameInfo.Type {
	case uModel.TicTacToeType:
		moveRepr = ticTacToeToMoveRepr(move)
	}
	return uModel.MoveInfo{
		PlayerID: gameInfo.Players[move.Player],
		MoveRepr: moveRepr,
	}
}

func ticTacToeToMoveRepr[T dModel.FigureType](move dModel.Move[T]) string {
	action := move.Actions[0]
	row := action.Args[0]
	col := action.Args[1]
	return fmt.Sprintf("%d %d", row, col)
}

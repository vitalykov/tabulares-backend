package mappers

import (
	"encoding/json"

	dModel "board-games/internal/delivery/model"
	uModel "board-games/internal/usecases/model"
)

func ToGameResponse(gameInfo *uModel.GameInfo) ([]byte, error) {
	output := dModel.GameResponse{
		Status: gameInfo.Status.String(),
		Turn:   gameInfo.Turn,
		Winner: gameInfo.Winner,
	}
	return json.Marshal(output)
}

func ToGameResponseWithMove(gameInfo *uModel.GameInfo, moveInfo uModel.MoveInfo) ([]byte, error) {
	// lastMove := gameInfo.Moves[len(gameInfo.Moves)-1]
	output := dModel.GameResponse{
		Status:   gameInfo.Status.String(),
		Turn:     gameInfo.Turn,
		Winner:   gameInfo.Winner,
		PlayerID: moveInfo.PlayerID,
		Move:     moveInfo.MoveRepr,
	}
	return json.Marshal(output)
}

func ToGameResponseFull(gameInfo *uModel.GameInfo) ([]byte, error) {
	output := dModel.GameResponseFull{
		ID:             gameInfo.ID,
		Name:           gameInfo.Type.String(),
		BoardWidth:     gameInfo.BoardWidth,
		BoardHeight:    gameInfo.BoardHeight,
		Players:        gameInfo.Players,
		Moves:          gameInfo.Moves,
		Winner:         gameInfo.Winner,
		Status:         gameInfo.Status.String(),
		Turn:           gameInfo.Turn,
		AdditionalInfo: gameInfo.AdditionalInfo,
	}
	return json.Marshal(output)
}

func ToNewGameResponse(gameInfo *uModel.GameInfo) ([]byte, error) {
	output := dModel.NewGameResponse{
		ID:             gameInfo.ID,
		Name:           gameInfo.Type.String(),
		Players:        gameInfo.Players,
		BoardWidth:     gameInfo.BoardWidth,
		BoardHeight:    gameInfo.BoardHeight,
		AdditionalInfo: gameInfo.AdditionalInfo,
	}
	return json.Marshal(output)
}

func ToMoveRequest(moveInfo uModel.MoveInfo) ([]byte, error) {
	output := dModel.MoveRequest{
		PlayerID: moveInfo.PlayerID,
		Move:     moveInfo.MoveRepr,
	}
	return json.Marshal(output)
}

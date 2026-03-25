package mappers

import (
	dModel "board-games/internal/delivery/model"
	uModel "board-games/internal/usecases/model"
	"encoding/json"
)

func GetMoveRequest(data []byte) (dModel.MoveRequest, error) {
	var moveInput dModel.MoveRequest
	err := json.Unmarshal(data, &moveInput)
	if err != nil {
		return dModel.MoveRequest{}, err
	}
	return moveInput, nil
}

func GetNewGameRequest(data []byte) (dModel.NewGameRequest, error) {
	var gameInput dModel.NewGameRequest
	err := json.Unmarshal(data, &gameInput)
	if err != nil {
		return dModel.NewGameRequest{}, err
	}
	return gameInput, nil
}

// func GetLoadGameInput(data []byte) (dModel.LoadGameInput, error) {
// 	var loadInput dModel.LoadGameInput
// 	err := json.Unmarshal(data, &loadInput)
// 	if err != nil {
// 		return dModel.LoadGameInput{}, err
// 	}
// 	return loadInput, nil
// }

// TODO: Are these mappers needed at all?

func ToMoveInfo(moveInput dModel.MoveRequest) uModel.MoveInfo {
	return uModel.MoveInfo{
		PlayerID: moveInput.PlayerID,
		MoveRepr: moveInput.Move,
	}
}

var gameTypes = map[string]uModel.GameType{
	"tic-tac-toe": uModel.TicTacToeType,
}

func ToNewGameInfo(input dModel.NewGameRequest) uModel.NewGameInfo {
	return uModel.NewGameInfo{
		Type:           gameTypes[input.Name],
		BoardWidth:     input.BoardWidth,
		BoardHeight:    input.BoardHeight,
		Players:        input.Players,
		AdditionalInfo: input.AdditionalInfo,
	}
}

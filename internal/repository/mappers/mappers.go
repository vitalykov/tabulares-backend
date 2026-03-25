package mappers

import (
	repoModel "board-games/internal/repository/model"
	ucModel "board-games/internal/usecases/model"
	"strconv"
	"strings"
)

const (
	PlayerDelimeter   = "@"
	MoveInfoDelimeter = "#"
)

var gameNameIDs = map[ucModel.GameType]int{
	ucModel.TicTacToeType: 1,
}

func ToGameData(gameInfo *ucModel.GameInfo) repoModel.GameData {
	return repoModel.GameData{
		ID:             gameInfo.ID,
		NameID:         gameNameIDs[gameInfo.Type],
		BoardWidth:     gameInfo.BoardWidth,
		BoardHeight:    gameInfo.BoardHeight,
		Players:        gameInfo.Players,
		Moves:          joinMoves(gameInfo.Moves),
		Winner:         gameInfo.Winner,
		AdditionalInfo: gameInfo.AdditionalInfo,
	}
}

func joinMoves(moves []ucModel.MoveInfo) string {
	var sb strings.Builder
	for _, mv := range moves {
		sb.WriteString(mv.PlayerID.String())
		sb.WriteString(PlayerDelimeter)
		sb.WriteString(mv.MoveRepr)
		sb.WriteString(MoveInfoDelimeter)
	}
	return sb.String()
}

var gameTypes = map[int]ucModel.GameType{
	1: ucModel.TicTacToeType,
}

func ToGameInfo(data repoModel.GameData) *ucModel.GameInfo {
	return &ucModel.GameInfo{
		ID:             data.ID,
		Type:           gameTypes[data.NameID],
		BoardWidth:     data.BoardWidth,
		BoardHeight:    data.BoardHeight,
		Players:        data.Players,
		Moves:          parseMoves(data.Moves),
		Winner:         data.Winner,
		AdditionalInfo: data.AdditionalInfo,
	}
}

// Parse string in format:
//
// "[playerID]@[MoveRepr]#..."
func parseMoves(s string) []ucModel.MoveInfo {
	moveReprs := strings.Split(s, MoveInfoDelimeter)
	moves := make([]ucModel.MoveInfo, len(moveReprs))
	for _, mvRepr := range moveReprs {
		l := strings.Split(mvRepr, PlayerDelimeter)
		playerID, _ := strconv.ParseInt(l[0], 10, 64)
		moves = append(moves, ucModel.MoveInfo{
			PlayerID: ucModel.PlayerID(playerID),
			MoveRepr: l[1],
		})
	}
	return moves
}

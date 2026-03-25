package boundaries

import "board-games/internal/usecases/model"

// Currently unused
type GameUsecases interface {
	CreateGame(gameInfo model.NewGameInfo) *model.GameInfo
	LoadGame(uuid model.UUID) *model.GameInfo
	StartGame(gameInfo *model.GameInfo)
	StopGame(gameInfo *model.GameInfo)
	CancelGame(gameInfo *model.GameInfo)
	MakeMove(gameInfo *model.GameInfo, moveInfo model.MoveInfo)
	UndoMove(gameInfo *model.GameInfo)
}

package handlers

import (
	"io"
	"log"
	"net/http"

	"board-games/internal/delivery/mappers"
	"board-games/internal/usecases/model"
	"board-games/internal/usecases/service"

	"github.com/google/uuid"
)

type httpHandler = func(w http.ResponseWriter, r *http.Request)

type Handlers interface {
	CreateGame() httpHandler
	LoadGame() httpHandler
	StartGame() httpHandler
	StopGame() httpHandler
	CancelGame() httpHandler
	MakeMove() httpHandler
	MakeAIMove() httpHandler
	UndoMove() httpHandler
	GetHint() httpHandler
}

type GameHandlers struct {
	gameMaster     *service.GameMaster
	gameInteractor *service.GameInteractorSwitch
}

func NewGameHandlers(gameMaster *service.GameMaster, gameInteractor *service.GameInteractorSwitch) *GameHandlers {
	return &GameHandlers{
		gameMaster:     gameMaster,
		gameInteractor: gameInteractor,
	}
}

func (h *GameHandlers) CreateGame() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: ", http.StatusBadRequest)
			return
		}
		input, err := mappers.GetNewGameRequest(data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: ", http.StatusBadRequest)
			return
		}
		info := mappers.ToNewGameInfo(input)
		gameInfo, err := h.gameMaster.CreateGame(info)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: ", http.StatusBadRequest)
			return
		}
		output, err := mappers.ToNewGameResponse(gameInfo)
		if err != nil {
			h.gameInteractor.CancelGame(gameInfo)
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(output)
		log.Println("Game created:", gameInfo.Type.String(), gameInfo.ID.String())
	}
}

func (h *GameHandlers) getGameInfo(id string) (*model.GameInfo, error) {
	gameID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	gameInfo, err := h.gameMaster.LoadGame(gameID)
	if err != nil {
		return nil, err
	}
	return gameInfo, nil
}

func (h *GameHandlers) StartGame() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if err = h.gameInteractor.StartGame(gameInfo); err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		output, err := mappers.ToGameResponse(gameInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		log.Println("Game started: ", gameInfo.Type.String(), gameInfo.ID.String())
	}
}

func (h *GameHandlers) LoadGame() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println("Load game:", err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if gameInfo.Status != model.Finished {
			gameInfo.Status = model.ReadyToStart
		}
		output, err := mappers.ToGameResponseFull(gameInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		log.Println("Game loaded:", gameInfo.Type.String(), gameInfo.ID.String())
	}
}

func (h *GameHandlers) StopGame() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if err = h.gameInteractor.StopGame(gameInfo); err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		log.Println("Game stopped:", gameInfo.Type.String(), gameInfo.ID.String())
	}
}

func (h *GameHandlers) CancelGame() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if err = h.gameInteractor.CancelGame(gameInfo); err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		log.Println("Game cancelled:", gameInfo.Type.String(), gameInfo.ID.String())
	}
}

func (h *GameHandlers) MakeMove() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if gameInfo.Status != model.InProgress {
			log.Println("Game is ", gameInfo.Status.String(), "but should be ", model.InProgress.String())
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: ", http.StatusBadRequest)
			return
		}
		input, err := mappers.GetMoveRequest(data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error: ", http.StatusBadRequest)
			return
		}
		moveInfo := mappers.ToMoveInfo(input)
		if err = h.gameInteractor.MakeMove(gameInfo, moveInfo); err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		output, err := mappers.ToGameResponse(gameInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		log.Println("Game:", gameInfo.ID.String(), "Player:", moveInfo.PlayerID, "Move:", moveInfo.MoveRepr)
	}
}

func (h *GameHandlers) MakeAIMove() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if gameInfo.Status != model.InProgress {
			log.Println("Game status:", gameInfo.Status.String(), "Expected:", model.InProgress.String())
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		moveInfo, err := h.gameInteractor.GetHint(gameInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if err = h.gameInteractor.MakeMove(gameInfo, moveInfo); err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		output, err := mappers.ToGameResponseWithMove(gameInfo, moveInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		log.Println("Game:", gameInfo.ID.String(), "Player:", moveInfo.PlayerID, "Move:", moveInfo.MoveRepr)
	}
}

func (h *GameHandlers) UndoMove() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		moveInfo, err := h.gameInteractor.UndoMove(gameInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		output, err := mappers.ToGameResponseWithMove(gameInfo, moveInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		log.Println("Undo move in game:", gameInfo.ID.String())
	}
}

func (h *GameHandlers) GetHint() httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		gameInfo, err := h.getGameInfo(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		if gameInfo.Status != model.InProgress {
			log.Println("Game status:", gameInfo.Status.String(), "Expected:", model.InProgress.String())
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		moveInfo, err := h.gameInteractor.GetHint(gameInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		output, err := mappers.ToMoveRequest(moveInfo)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error:", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		log.Println("Game:", gameInfo.ID.String(), "Player:", moveInfo.PlayerID, "Move:", moveInfo.MoveRepr)
	}
}

package routes

import (
	"board-games/internal/delivery/handlers"
	"net/http"
)

func NewRouter(h *handlers.GameHandlers) *http.ServeMux {
	router := http.NewServeMux()
	MapHandlers(router, h)
	return router
}

const GamePath = "/game/"

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func makeGamePath(method, path string) string {
	return method + " " + GamePath + path
}

func MapHandlers(mux *http.ServeMux, h *handlers.GameHandlers) {
	mux.HandleFunc(makeGamePath(POST, "create/"), h.CreateGame())
	mux.HandleFunc(makeGamePath(GET, "load/{id}/"), h.LoadGame())
	mux.HandleFunc(makeGamePath(PUT, "start/{id}/"), h.StartGame())
	mux.HandleFunc(makeGamePath(PUT, "stop/{id}/"), h.StopGame())
	mux.HandleFunc(makeGamePath(DELETE, "cancel/{id}/"), h.CancelGame())
	mux.HandleFunc(makeGamePath(POST, "move/{id}/"), h.MakeMove())
	mux.HandleFunc(makeGamePath(POST, "ai_move/{id}/"), h.MakeAIMove())
	mux.HandleFunc(makeGamePath(PUT, "undo/{id}/"), h.UndoMove())
	mux.HandleFunc(makeGamePath(GET, "hint/{id}/"), h.GetHint())
}

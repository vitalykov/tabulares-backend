package routes

import (
	"net/http"

	"board-games/internal/delivery/handlers"
	"board-games/pkg/web"
)

func NewRouter(h *handlers.GameHandlers) *http.ServeMux {
	router := http.NewServeMux()
	MapHandlers(router, h)
	return router
}

const gamePath = "game"

func MapHandlers(mux *http.ServeMux, h *handlers.GameHandlers) {
	mux.HandleFunc(web.MakePath(web.POST, gamePath, "create"), h.CreateGame())
	mux.HandleFunc(web.MakePath(web.GET, gamePath, "load/{id}"), h.LoadGame())
	mux.HandleFunc(web.MakePath(web.PUT, gamePath, "start/{id}"), h.StartGame())
	mux.HandleFunc(web.MakePath(web.PUT, gamePath, "stop/{id}"), h.StopGame())
	mux.HandleFunc(web.MakePath(web.DELETE, gamePath, "cancel/{id}"), h.CancelGame())
	mux.HandleFunc(web.MakePath(web.POST, gamePath, "move/{id}"), h.MakeMove())
	mux.HandleFunc(web.MakePath(web.POST, gamePath, "ai_move/{id}"), h.MakeAIMove())
	mux.HandleFunc(web.MakePath(web.PUT, gamePath, "undo/{id}"), h.UndoMove())
	mux.HandleFunc(web.MakePath(web.GET, gamePath, "hint/{id}"), h.GetHint())
}

package routes

import (
	"board-games/pkg/web"
	"net/http"

	"board-games/internal/auth"
)

func NewRouter(h auth.Handlers) *http.ServeMux {
	router := http.NewServeMux()
	mapHandlers(router, h)
	return router
}

const authBase = "auth"

func mapHandlers(mux *http.ServeMux, h auth.Handlers) {
	mux.HandleFunc(web.MakePath(web.POST, authBase, "register"), h.Register())
	mux.HandleFunc(web.MakePath(web.POST, authBase, "login"), h.Login())
	mux.HandleFunc(web.MakePath(web.POST, authBase, "logout"), h.Logout())
	mux.HandleFunc(web.MakePath(web.GET, authBase, "{id}"), h.GetByID())
	mux.HandleFunc(web.MakePath(web.GET, authBase, "find/{username}"), h.FindByUsername())
	mux.HandleFunc(web.MakePath(web.DELETE, authBase, "{id}"), h.Delete())
}

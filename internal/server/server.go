package server

import (
	"context"
	"net/http"
	"time"

	"board-games/internal/delivery/handlers"
	"board-games/internal/delivery/routes"
	"board-games/internal/repository/cache"
	"board-games/internal/usecases/boundaries"
	"board-games/internal/usecases/service"

	"go.uber.org/fx"
)

const addr = ":8080"

type HTTPGameServer struct {
	server http.Server
}

func NewHTTPGameServer(router *http.ServeMux) *HTTPGameServer {
	return &HTTPGameServer{
		server: http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *HTTPGameServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPGameServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func CreateServer() fx.Option {
	return fx.Options(
		fx.Provide(
			NewHTTPGameServer,
			routes.NewRouter,
			fx.Annotate(
				cache.NewGameCache,
				fx.As(new(boundaries.GameRepository)),
			),
			service.NewGameMaster,
			service.NewGameInteractorSwitch,
			handlers.NewGameHandlers,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *HTTPGameServer) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go server.Run()
						return nil
					},
					OnStop: server.Shutdown,
				})
			},
		),
	)
}

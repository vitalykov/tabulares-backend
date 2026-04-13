package server

import (
	"context"
	"net/http"
	"time"

	"board-games/config"
	"board-games/internal/auth"
	authHTTP "board-games/internal/auth/delivery/http"
	"board-games/internal/auth/delivery/routes"
	"board-games/internal/auth/repository/pg"
	"board-games/internal/auth/usecase/service"
	"board-games/pkg/db"

	"go.uber.org/fx"
)

const addr = ":8081"

type HTTPAuthServer struct {
	server http.Server
}

func NewHTTPAuthServer(router *http.ServeMux) *HTTPAuthServer {
	return &HTTPAuthServer{
		server: http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *HTTPAuthServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPAuthServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func CreateServer() fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewDBConfig,
			NewHTTPAuthServer,
			routes.NewRouter,
			service.NewUserManager,
			db.NewPostgresPool,
			fx.Annotate(
				pg.NewPgUserRepo,
				fx.As(new(auth.UserRepo)),
			),
			fx.Annotate(
				authHTTP.NewAuthHandlers,
				fx.As(new(auth.Handlers)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, server *HTTPAuthServer) {
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

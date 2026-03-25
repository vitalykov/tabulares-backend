package main

import (
	"board-games/internal/server"

	"go.uber.org/fx"
)

func main() {
	// s := server.NewHTTPGameServer(8080)
	// s.Run()
	fx.New(server.CreateServer()).Run()
}

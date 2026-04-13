package main

import (
	"board-games/internal/auth/server"

	"go.uber.org/fx"
)

func main() {
	fx.New(server.CreateServer()).Run()
}

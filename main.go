package main

import (
	"chat-ws/internal/ws"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			ws.NewRoom,
			ws.NewMux,
		),

		fx.Invoke(
			ws.StartServer,
		),
	)

	app.Run()

}

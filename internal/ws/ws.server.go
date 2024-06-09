package ws

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func NewMux(r *room) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/room", r)
	return mux
}

func StartServer(lc fx.Lifecycle, mux *http.ServeMux) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := ":8080"
			log.Println("Starting web server on", addr)
			go func() {
				if err := http.ListenAndServe(addr, mux); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down the server.")
			return nil
		},
	})
}

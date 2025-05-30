package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ksamirdev/lyrica/env"
	"github.com/ksamirdev/lyrica/web/routes"
)

type Application struct {
	config env.Config
}

func (app *Application) Serve() error {
	port := app.config.PORT

	log.Printf("🚀 Server listening to port %s", port)

	srv := &http.Server{
		Addr:    net.JoinHostPort("localhost", port),
		Handler: routes.Routes(),
	}

	return srv.ListenAndServe()

}

func init() {
	env.Load()
}

func main() {
	app := Application{
		config: env.DefaultConfig,
	}

	if err := app.Serve(); err != nil {
		log.Fatalf("Error starting server: %s\b", err.Error())
	}
}

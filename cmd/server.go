package cmd

import (
	"log"
	"net/http"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func (app *Application) Run(mux *http.ServeMux) {

	server := http.Server{
		Addr:    app.Config.Addr,
		Handler: mux,
	}
	log.Println("Application runnning on", app.Config.Addr)
	server.ListenAndServe()
}

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
	mux  *http.ServeMux
}

func (app *Application) Mount(mux *http.ServeMux) {
	app.Config.mux = mux
}

func (app *Application) Run() {
	server := http.Server{
		Addr:    app.Config.Addr,
		Handler: app.Config.mux,
	}

	log.Println("Application runnning on", app.Config.Addr)
	server.ListenAndServe()
}

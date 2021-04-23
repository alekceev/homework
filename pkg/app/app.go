package app

import (
	"homework/pkg/interfaces"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     interfaces.DB
}

func NewApp() *App {
	app := &App{Router: mux.NewRouter()}
	return app
}

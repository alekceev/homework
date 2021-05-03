package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"homework/pkg/controllers/index"
	"homework/pkg/controllers/items"
	"homework/pkg/interfaces"
	"homework/pkg/repositories"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     interfaces.DB
}

func NewApp() *App {
	app := &App{Router: mux.NewRouter()}
	return app
}

func (app *App) Start(host, port string) error {
	if app.DB == nil {
		panic("nil db")
	}

	// Repositories
	itemsRepo := repositories.NewItemRepository(app.DB)

	// Item Controller
	itemsCtrl := items.NewItemsController(itemsRepo)

	for _, route := range itemsCtrl.Routes() {
		app.Router.HandleFunc(route.Route, route.Handler).Methods(route.Methods...)
	}

	// Index Controller
	indexCtr := index.NewIndexController()

	for _, route := range indexCtr.Routes() {
		app.Router.HandleFunc(route.Route, route.Handler).Methods(route.Methods...)
	}

	fmt.Printf("Server is listening http://%s:%s ...", host, port)
	return http.ListenAndServe(host+":"+port, app.Router)
}

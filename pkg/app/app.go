package app

import (
	"homework/pkg/controllers/index"
	"homework/pkg/controllers/items"
	"homework/pkg/database"
	"homework/pkg/interfaces"
	"homework/pkg/repositories"
	server "homework/pkg/services/Server"
	"homework/pkg/services/discounter"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	DB       interfaces.DB
	Config   *Config
	Services []interfaces.Service
}

func NewApp() *App {
	app := &App{
		Config:   ParseConfig(),
		Router:   mux.NewRouter(),
		Services: make([]interfaces.Service, 0),
	}

	// init DB
	var err error
	app.DB, err = database.Connect(app.Config.DbHost)
	if err != nil {
		log.Fatalf("Db error: %v", err)
		os.Exit(1)
	}

	return app
}

func (app *App) Close() {

	for _, service := range app.Services {
		service.Stop()
	}

	app.DB.Close()

	log.Println("End...")
}

func (app *App) Start() error {
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

	http.Handle("/", app.Router)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web"))))

	// web server
	app.Services = append(app.Services,
		server.NewWebServer(
			&http.Server{
				Addr: app.Config.Host + ":" + app.Config.Port,
			},
		),
	)

	//over services
	app.Services = append(app.Services, discounter.NewDiscountService(itemsRepo, app.Config.DiscountUrl))

	// start Services
	app.StartServices()

	return nil
}

func (app *App) StartServices() {
	var wg sync.WaitGroup
	for _, service := range app.Services {
		wg.Add(1)
		go service.Start()
	}
	wg.Wait()
}

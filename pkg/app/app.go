package app

import (
	"context"
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
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	DB       interfaces.DB
	Config   *Config
	Services []interfaces.Service
	ctx      context.Context
}

func NewApp() *App {
	app := &App{
		Config:   ParseConfig(),
		Router:   mux.NewRouter(),
		Services: make([]interfaces.Service, 0),
		ctx:      context.Background(),
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

func (app *App) AddService(service interfaces.Service) {
	app.Services = append(app.Services, service)
}

func (app *App) Stop() {
	ctx, cancel := context.WithTimeout(app.ctx, 15*time.Second)
	defer cancel()

	for _, service := range app.Services {
		service.Stop(ctx)
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

	// Index Controller
	indexCtr := index.NewIndexController()

	for _, route := range indexCtr.Routes() {
		app.Router.HandleFunc(route.Route, route.Handler).Methods(route.Methods...)
	}

	// Item Controller
	itemsCtrl := items.NewItemsController(itemsRepo)

	for _, route := range itemsCtrl.Routes() {
		app.Router.HandleFunc(route.Route, route.Handler).Methods(route.Methods...)
	}

	http.Handle("/", app.Router)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web"))))

	// web server
	app.AddService(
		server.NewWebServer(
			&http.Server{
				Addr: app.Config.Host + ":" + app.Config.Port,
			},
		),
	)

	//other services
	app.AddService(discounter.NewDiscountService(itemsRepo, app.Config.DiscountUrl))

	// start Services
	app.StartServices()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	app.Stop()

	return nil
}

func (app *App) StartServices() {
	for _, service := range app.Services {
		go service.Start(app.ctx)
	}
}

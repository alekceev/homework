package app

import (
	"fmt"
	"homework/pkg/controller"
	"homework/pkg/repositories"
	"net/http"
)

func (app *App) InitRouters() {

	itemRepo := repositories.NewItemRepository(app.DB.Raw())

	// Item Handler
	item := controller.NewItemHandler(itemRepo)
	app.Router.HandleFunc("/items", item.ItemsHandler).Methods(http.MethodGet, http.MethodPost)
	app.Router.HandleFunc("/items/{id:[0-9]+}", item.ItemHandler).Methods(http.MethodGet, http.MethodDelete)

	// index
	app.Router.HandleFunc("/", app.IndexHandler()).Methods(http.MethodGet)
}

func (app *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			Create item
			curl -sL -XPOST http://localhost:8081/items -d '{"name":"Intel Core I3", "description":"Intel Cole i3", "article":"i3-e100", "category":"Процессоры", "price": 5}'
	
			Get All items
			curl -sL http://localhost:8081/items
	
			Get item 1
			curl -sL http://localhost:8081/items/1
	
			Delete item 1
			curl -sL -XDELETE http://localhost:8081/items/1
		`)
	}
}

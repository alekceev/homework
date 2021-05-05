package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"homework/pkg/repositories"
	"homework/pkg/server"
	"homework/pkg/sqlite"

	"github.com/gorilla/mux"
	"github.com/namsral/flag"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		Create item
		curl -sL -XPOST http://localhost:8081/items -d '{"name":"Cake", "description":"Delicious Cake", "price": 5}'

		Get All items
		curl -sL http://localhost:8081/items

		Get item 1
		curl -sL http://localhost:8081/items/1

		Delete item 1
		curl -sL -XDELETE http://localhost:8081/items/1
	`)
}

func main() {

	var port string
	flag.StringVar(&port, "port", "8081", "Server port")
	flag.Parse()

	// init Db
	db, err := sqlite.SqliteInit()
	if err != nil {
		log.Fatalf("Db error: %v", err)
		os.Exit(1)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Db close error: %v", err)
		}
	}()

	itemRepo := repositories.NewItemRepository(db)

	// Item Handler
	item := server.NewItemHandler(itemRepo)

	router := mux.NewRouter()
	router.HandleFunc("/items", item.ItemsHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/items/{id:[0-9]+}", item.ItemHandler).Methods(http.MethodGet, http.MethodDelete)
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening http://localhost:" + port + "...")
	http.ListenAndServe(":"+port, nil)
}

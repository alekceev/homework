package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"homework/pkg/globals"
	"homework/pkg/items"
	"homework/pkg/sqlite"

	"github.com/gorilla/mux"
)

const port = "8081"

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		// Get items
		items, err := items.GetAll()
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Items:\n\t%#v\n\n", items)

		data, err := json.Marshal(items)
		if err != nil {
			panic(err)
		}
		renderJSON(w, data)

	case http.MethodPost:
		// Create item
		log.Println("create item")

		decoder := json.NewDecoder(r.Body)
		var item items.Item
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		log.Println(item.Name)

		item, err = items.CreateItem(item)
		if err != nil {
			panic(err)
		}

		data, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		renderJSON(w, data)

	default:
		// Give an error message.
		renderJSON(w, []byte(`{"error": 1}`))
	}
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	switch r.Method {
	case http.MethodGet:
		// Get item
		item, err := items.GetItem(id)
		if err != nil {
			log.Printf("Error get item: %d %v", id, err)
			renderJSON(w, []byte(`{"error":"Not found"}`))
			return
		}

		data, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		renderJSON(w, data)

	case http.MethodDelete:
		// Remove item
		err := items.DeleteItem(id)
		if err != nil {
			panic(err)
		}
		renderJSON(w, []byte(`{"ok": 1}`))

	default:
		// Give an error message.
		renderJSON(w, []byte(`{"error": 1}`))
	}
}

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

	// init Db
	var err error
	globals.Db, err = sqlite.SqliteInit()
	if err != nil {
		log.Fatalf("Db error: %v", err)
		os.Exit(1)
	}
	defer func() {
		err := globals.Db.Close()
		if err != nil {
			log.Fatalf("Db close error: %v", err)
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/items", itemsHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/items/{id:[0-9]+}", itemHandler).Methods(http.MethodGet, http.MethodDelete)
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening http://localhost:" + port + "...")
	http.ListenAndServe(":"+port, nil)
}

func renderJSON(w http.ResponseWriter, data []byte) {
	log.Printf("response:\n\t%s\n", string(data))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

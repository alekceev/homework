package server

import (
	"encoding/json"
	"homework/pkg/interfaces"
	"homework/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type itemHandler struct {
	repo interfaces.Repository
}

func NewItemHandler(repo interfaces.Repository) *itemHandler {
	return &itemHandler{repo: repo}
}

func (h *itemHandler) ItemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		// Get items
		items, err := h.repo.GetAll()
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
		var item models.Item
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		log.Println(item.Name)

		err = h.repo.Save(&item)
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

func (h *itemHandler) ItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	switch r.Method {
	case http.MethodGet:
		// Get item
		item, err := h.repo.Get(id)
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
		err := h.repo.Delete(id)
		if err != nil {
			panic(err)
		}
		renderJSON(w, []byte(`{"ok": 1}`))

	default:
		// Give an error message.
		renderJSON(w, []byte(`{"error": 1}`))
	}
}

func renderJSON(w http.ResponseWriter, data []byte) {
	log.Printf("response:\n\t%s\n", string(data))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

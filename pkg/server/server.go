package server

import (
	"encoding/json"
	"fmt"
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
			log.Printf("Error GetAll: %v", err)
			renderJSON(w, `{"error":"Error get items"}`)
			return
		}

		renderJSON(w, items)

	case http.MethodPost:
		// Create item
		log.Println("create item")

		decoder := json.NewDecoder(r.Body)
		var item models.Item
		err := decoder.Decode(&item)
		if err != nil {
			log.Printf("Error decode: %v", err)
			renderJSON(w, `{"error":"Error"}`)
			return
		}

		err = h.repo.Save(&item)
		if err != nil {
			log.Printf("Error save item: %v", err)
			renderJSON(w, `{"error":"Internal error"}`)
			return
		}

		renderJSON(w, item)

	default:
		// Give an error message.
		log.Println("Not found")
		renderJSON(w, `{"error": "Not found"}`)
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
			renderJSON(w, `{"error":"Not found"}`)
			return
		}

		renderJSON(w, item)

	case http.MethodDelete:
		// Remove item
		err := h.repo.Delete(id)
		if err != nil {
			log.Printf("Error delete item: %d %v", id, err)
			renderJSON(w, `{"error":"Not found"}`)
			return
		}
		renderJSON(w, `{"ok": "success"}`)

	default:
		// Give an error message.
		renderJSON(w, []byte(`{"error": "Not found"}`))
	}
}

func renderJSON(w http.ResponseWriter, data interface{}) {
	status := http.StatusOK

	var (
		resp []byte
		err  error
	)

	if data == nil {
		status = http.StatusNoContent
	} else {

		switch data := data.(type) {
		case string:
			resp = []byte(data)
		case []byte:
			resp = data
		default:
			resp, err = json.Marshal(data)
			if err != nil {
				status = http.StatusInternalServerError
				resp = []byte(fmt.Sprintf(`{"error":"%v"}`, err))
			}
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(resp)
}

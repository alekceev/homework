package items

import (
	"encoding/json"
	"homework/pkg/controllers/routes"
	"homework/pkg/interfaces"
	"homework/pkg/models"
	"homework/pkg/renderer"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ItemsController struct {
	repo interfaces.Repository
}

func NewItemsController(repo interfaces.Repository) *ItemsController {
	return &ItemsController{repo: repo}
}

func (h *ItemsController) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)

	renderer := renderer.NewRender(r)
	templates := []string{
		"./pkg/templates/layout.tmpl",
		"./pkg/templates/items/items.tmpl",
	}

	switch r.Method {
	case http.MethodGet:
		// Get items
		items, err := h.repo.GetAll()
		if err != nil {
			log.Printf("Error GetAll: %v", err)
			renderer.Render(w, `{"error":"Error get items"}`, templates, http.StatusOK)
			return
		}

		renderer.Render(w, items, templates, http.StatusOK)

	case http.MethodPost:
		// Create item
		log.Println("create item")

		decoder := json.NewDecoder(r.Body)
		var item models.Item
		err := decoder.Decode(&item)
		if err != nil {
			log.Printf("Error decode: %v", err)
			renderer.Render(w, `{"error":"Error"}`, templates, http.StatusOK)
			return
		}

		err = h.repo.Save(&item)
		if err != nil {
			log.Printf("Error save item: %v", err)
			renderer.Render(w, `{"error":"Internal error"}`, templates, http.StatusOK)
			return
		}

		renderer.Render(w, item, templates, http.StatusOK)

	default:
		// Give an error message.
		log.Println("Not found")
		renderer.Render(w, `{"error": "Not found"}`, templates, http.StatusOK)
	}
}

func (h *ItemsController) Get(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)

	renderer := renderer.NewRender(r)
	templates := []string{
		"./pkg/templates/layout.tmpl",
		"./pkg/templates/items/item.tmpl",
	}

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	switch r.Method {
	case http.MethodGet:
		// Get item
		item, err := h.repo.Get(id)
		if err != nil {
			log.Printf("Error get item: %d %v", id, err)
			renderer.Render(w, `{"error":"Not found"}`, templates, http.StatusOK)
			return
		}

		renderer.Render(w, item, templates, http.StatusOK)

	case http.MethodDelete:
		// Remove item
		err := h.repo.Delete(id)
		if err != nil {
			log.Printf("Error delete item: %d %v", id, err)
			renderer.Render(w, `{"error":"Not found"}`, templates, http.StatusOK)
			return
		}
		renderer.Render(w, `{"ok": "success"}`, templates, http.StatusOK)

	default:
		// Give an error message.
		renderer.Render(w, []byte(`{"error": "Not found"}`), templates, http.StatusOK)
	}
}

func (h *ItemsController) Routes() []routes.Route {
	return []routes.Route{
		{Route: "/items", Handler: h.GetAll, Methods: []string{http.MethodGet, http.MethodPost}},
		{Route: "/items/{id:[0-9]+}", Handler: h.Get, Methods: []string{http.MethodGet, http.MethodDelete}},
	}
}

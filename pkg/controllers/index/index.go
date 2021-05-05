package index

import (
	"fmt"
	"homework/pkg/controllers/routes"
	"net/http"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (h *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	// todo: render in json (or Renders)
	fmt.Fprint(w, `
			Create item
			curl -sL -XPOST http://localhost:8081/items -d '{"name":"Intel Core I3", "description":"Intel Cole i3", "number":"i3-e100", "category":"Процессоры", "price": 5}'
	
			Get All items
			curl -sL http://localhost:8081/items
	
			Get item 1
			curl -sL http://localhost:8081/items/1
	
			Delete item 1
			curl -sL -XDELETE http://localhost:8081/items/1
		`)
}

func (h *IndexController) Routes() []routes.Route {
	return []routes.Route{
		{Route: "/", Handler: h.Index, Methods: []string{http.MethodGet}},
	}
}

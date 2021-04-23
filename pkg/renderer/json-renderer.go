package renderer

import (
	"encoding/json"
	"fmt"
	"homework/pkg/interfaces"
	"net/http"
)

type JsonRenderer struct {
	status  int
	headers map[string]string
}

var _ interfaces.Renderer = &JsonRenderer{}

func NewJsonRenderer() *JsonRenderer {
	return &JsonRenderer{
		headers: map[string]string{"Content-Type": "application/json; charset=utf-8"},
	}
}

func (r *JsonRenderer) Render(data interface{}, _ []string, status int) []byte {
	r.status = status
	if data == nil && status == http.StatusOK {
		r.status = http.StatusNoContent
		return nil
	} else {

		switch data := data.(type) {
		case string:
			return []byte(data)
		case []byte:
			return data
		default:
			resp, err := json.Marshal(data)
			if err != nil {
				r.status = http.StatusInternalServerError
				return []byte(fmt.Sprintf(`{"error":"%v"}`, err))
			}
			return resp
		}
	}
}

func (r *JsonRenderer) Status() int {
	return r.status
}

func (r *JsonRenderer) Headers() map[string]string {
	return r.headers
}

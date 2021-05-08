package renderer

import (
	"homework/pkg/interfaces"
	"log"
	"net/http"
)

type Render struct {
	req      *http.Request
	renderer interfaces.Renderer
}

func NewRender(r *http.Request) *Render {
	// log.Printf("RENDERER: %s %s %#v\n", r.Method, r.URL.Path, r.Header)

	params := r.URL.Query()
	format := params.Get("format")

	if len(format) == 0 {
		switch r.Header.Get("Content-Type") {
		case "application/json":
			format = "json"
		case "text/html":
			format = "html"
		default:
			format = "json"
		}
	}

	var renderer interfaces.Renderer

	switch format {
	case "json":
		renderer = NewJsonRenderer()
	case "html":
		renderer = NewHtmlRenderer()
	default:
		renderer = NewJsonRenderer()
	}

	log.Printf("%s %#v\n", format, renderer)

	return &Render{req: r, renderer: renderer}
}

func (r *Render) Render(w http.ResponseWriter, data interface{}, templates []string, status int) {
	resp := r.renderer.Render(data, templates, status)
	for k, v := range r.renderer.Headers() {
		w.Header().Add(k, v)
	}
	w.WriteHeader(r.renderer.Status())
	w.Write(resp)
}

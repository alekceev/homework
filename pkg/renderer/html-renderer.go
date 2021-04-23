package renderer

import (
	"bytes"
	"fmt"
	"homework/pkg/interfaces"
	"log"
	"net/http"
	"text/template"
)

type HtmlRenderer struct {
	status  int
	headers map[string]string
}

var _ interfaces.Renderer = &HtmlRenderer{}

func NewHtmlRenderer() *HtmlRenderer {
	return &HtmlRenderer{
		headers: map[string]string{"Content-Type": "text/html; charset=utf-8"},
	}
}

func (r *HtmlRenderer) Render(data interface{}, templates []string, status int) []byte {
	r.status = status

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err.Error())
		r.status = http.StatusInternalServerError
		return []byte(fmt.Sprintf(`"error":"%v"`, err))
	}

	buf := new(bytes.Buffer)
	defer buf.Reset()

	if err := ts.Execute(buf, data); err != nil {
		r.status = http.StatusInternalServerError
		return []byte(fmt.Sprintf(`"error":"%v"`, err))
	}
	return buf.Bytes()
}

func (r *HtmlRenderer) Status() int {
	return r.status
}

func (r *HtmlRenderer) Headers() map[string]string {
	return r.headers
}

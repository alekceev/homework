package routes

import "net/http"

type Route struct {
	Route   string
	Handler http.HandlerFunc
	Methods []string
}

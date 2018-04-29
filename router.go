package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new mux router
// and adds the routes to itself
// it also wraps itself in a Logger function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

package service

import "github.com/gorilla/mux"

// NewRouter return a pointer to a mux.Router we can use as a handler
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		// Attach each route
		// Use a Builder-like pattern to set each route up
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

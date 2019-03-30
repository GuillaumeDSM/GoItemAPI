package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	loadItems("vdms.json")
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

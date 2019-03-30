package routing

import (
	"net/http"

	"github.com/GuillaumeDSM/GoItemAPI/model"
	"github.com/GuillaumeDSM/GoItemAPI/util"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	model.LoadItems("vdms.json")
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = util.Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

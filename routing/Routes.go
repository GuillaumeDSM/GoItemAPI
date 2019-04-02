package routing

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ItemIndex",
		"GET",
		"/items",
		ItemList,
	},
	Route{
		"ItemIndex",
		"GET",
		"/item",
		ItemIndex,
	},
	Route{
		"ItemIndex",
		"POST",
		"/item",
		ItemCreate,
	},
	Route{
		"ItemShow",
		"GET",
		"/item/{ItemId}",
		ItemShow,
	},
	Route{
		"ItemApprove",
		"POST",
		"/item/{ItemId}/approve",
		ItemApprove,
	},
}

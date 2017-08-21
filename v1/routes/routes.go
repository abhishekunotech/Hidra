package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Unotechsoftware/Hydra/v1/handlers"
    "github.com/Unotechsoftware/Hydra/v1/utils"

)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

    	var handler http.Handler
	handler = route.HandlerFunc
	handler = utils.Logger(handler, route.Name)


		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"getTicketDetails",
		"GET",
		"/getTicketDetails",
		handlers.GetTicketDetails,
	},
	Route{
		"getCILogs",
		"GET",
		"/getCILogs",
		handlers.GetCILogs,
	},
	Route{
		"getCIJobs",
		"GET",
		"/getCIJobs",
		handlers.GetCIJobs,
	},
	Route{
		"getCIDetails",
		"GET",
		"/getCIDetails",
		handlers.GetCIDetails,
	},
	Route{
		"Ticketcreate",
		"POST",
		"/Ticketcreate",
		handlers.Ticketcreate,
	},
}

/*
	The package Routes defines properties of an HTTP endpoint.
	At runtime, the router will associate each Route with a http.Handler object, and use the Route properties to determine which Handler should be invoked.
	Basically Routes will define routes for the different functions.
	Install using go install in this directory.

	Author: Operations Management Team - Unotech Software.
*/
package routes

import (
	"github.com/Unotechsoftware/Hydra/handlers"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	// Name is a key specifying which HTTP handler the router should associate
	Name string
	//Method is any valid HTTP method
	Method string
	//Pattern contains a path pattern
	Pattern string
	//handler
	HandlerFunc http.HandlerFunc
}

//Routes is a Route collection.
type Routes []Route

func NewRouter() *mux.Router {
	//Create a new mux router for given handler
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		//Create the Route for the requested method,name,pattern and handler.

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

//Create different routes for required functions using name, method, path pattern and handler.
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
		"getListOfWorkOrders",
		"GET",
		"/getListOfWorkOrders",
		handlers.GetListOfWorkOrders,
	},
	Route{
		"Ticketcreate",
		"POST",
		"/Ticketcreate",
		handlers.Ticketcreate,
	},
	Route{
		"getLinkedChange",
		"GET",
		"/getLinkedChange",
		handlers.GetLinkedChange,
	},
	Route{

		"getListOfLinkedChange",
		"GET",
		"/getListOfLinkedChange",
		handlers.GetListOfLinkedChange,
	},
        Route{
		"TicketAll",
		"GET",
		"/TicketAll",
		handlers.TicketAll,

	},
	Route{
		"GetCountOfOpenTickets",
		"GET",
		"/getCountOfOpenTicketsCustomerUser",
		handlers.GetCountOfOpenTicketsCustomerUser,
	},
	Route{
		"getLinkedTickets",
		"GET",
		"/felicity/getLinkedTickets",
		handlers.GetLinkedTickets,
	},
	Route{
		"felicityLogin",
		"GET",
		"/felicity/isValidUser",
		handlers.IsValidFelicityUser,
	},
	Route{

		"getCustomerInfobyTicketID",
		"GET",
		"/felicity/getCustomerInfobyTicketID",
		handlers.GetCustomerInfobyTicketID,
	},
	Route{

                "getListOfCIs",
                "GET",
                "/getListOfCIs",
                handlers.GetListOfCIs,

        },
	Route{
                "getTicketInfo",
                "GET",
                "/getTicketInfo",
                handlers.GetTicketInfo,
        },


	Route{
		"getCountClosedTickets",
		"GET",
		"/felicity/getCountClosedTickets",
		handlers.GetCountofClosedTickets,
	},
	Route{
		"getListofWorkorderGraph",
		"GET",
		"/felicity/getListofWorkorderGraph",
		handlers.GetListofWorkorderGraph,
	},
	Route{
		"setUserColumnPreferences",
		"POST",
		"/felicity/setUserColumnPreferences",
		handlers.SetUserColumnPreferences,
	},
	Route{
		"getUserColumnPreferences",
		"GET",
		"/getUserColumnPreference",
		handlers.GetUserColumnPreferences,
	},
	Route{
		"getListAssignedQueue",
		"GET",
		"/getListAssignedQueue",
		handlers.GetListAssignedQueue,
	},
	Route{
		"getListAgents",
		"GET",
		"/getListAgents",
		handlers.GetListAgents,
	},
	Route{
		"getListGroupFilter",
		"GET",
		"/getListGroupFilter",
		handlers.GetListGroupFilter,
	},
	Route{
		"getListTicketState",
		"GET",
		"/getListTicketState",
		handlers.GetListTicketState,
	},
	Route{
		"listPriority",
		"GET",
		"/listPriority",
		handlers.ListPriority,
	},
	Route{
		"custListAjax",
		"GET",
		"/custListAjax",
		handlers.CustListAjax,
	},
	Route{
		"requestListAjax",
		"GET",
		"/requestListAjax",
		handlers.RequestListAjax,
	},
}

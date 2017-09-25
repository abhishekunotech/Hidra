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
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
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
	PopulateRoutes()
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


func PopulateRoutes() {
	ConfObj := lerna.ReturnConfigObject()
	RouteMapString := ConfObj.GetStringMap("routes")
	RouteKeyArray := lerna.GetKeyArray(RouteMapString)
	for _, routeVal := range RouteKeyArray {
		var tempRouteObj Route
		var tempHandler handlers.Handler
		tempRouteObj.Name = routeVal
		tempRouteObj.Method = ConfObj.GetString("routes." + routeVal + ".method")
		tempRouteObj.Pattern = ConfObj.GetString("routes." + routeVal + ".URI")

		tempRouteObj.HandlerFunc = http.HandlerFunc(reflect.ValueOf(&tempHandler).MethodByName(ConfObj.GetString("routes."+routeVal+".handler")).Interface().(func(w http.ResponseWriter, r *http.Request)))
		routes = append(routes, tempRouteObj)
	}
}

var routes Routes


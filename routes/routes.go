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

// As defined in the other Route Structure
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


// Creates a New Router Object
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

// Uses Lerna to Populate the Routes Global Object with the Routes from the Configuration
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

// A mapping between the Name and the URI of the Route
type Route_Name_URI struct{
	Name	string
	URI	string
}

// A SLice of Route_Name_URI Structure
type Exposed_routes []Route_Name_URI

// Exports the Names and URIs of API Exposed through configuration
func GetExposedRouteList() Exposed_routes{
	var ExposedRoutes Exposed_routes
	for _,vals := range routes{
		var temp Route_Name_URI
		temp.Name = vals.Name
		temp.URI = vals.Pattern
		ExposedRoutes = append(ExposedRoutes,temp)
	}
	return ExposedRoutes
}

// Exports the names of the API exposed through Configuration
func GetRouteNames() []string{
	var ReturnVal []string
	for _,Names := range routes{
		ReturnVal = append(ReturnVal,Names.Name)
	}
	return ReturnVal
}

// Exports the URIs of the API exposed through Configuration
func GetRouteURIs() []string{
	var ReturnVal []string
	for _,URIs := range routes{
		ReturnVal = append(ReturnVal,URIs.Pattern)
	}
	return ReturnVal
}

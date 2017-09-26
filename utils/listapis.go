package utils

import(
	"github.com/Unotechsoftware/Hydra/lerna"
	"github.com/Unotechsoftware/Hydra/models"
)

func ListRoutes() []models.APIObject{
	var routes []models.APIObject
        ConfObj := lerna.ReturnConfigObject()
        RouteMapString := ConfObj.GetStringMap("routes")
        RouteKeyArray := lerna.GetKeyArray(RouteMapString)
        for _, routeVal := range RouteKeyArray {
                var tempRouteObj models.APIObject
                tempRouteObj.Name = routeVal
		tempRouteObj.Method = ConfObj.GetString("routes." + routeVal + ".method")
                tempRouteObj.URI = ConfObj.GetString("routes." + routeVal + ".URI")

                routes = append(routes, tempRouteObj)
        }
	return routes
}


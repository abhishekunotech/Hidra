package models

import(
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


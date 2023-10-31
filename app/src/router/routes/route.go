package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Config(router *mux.Router) {
	var routes []Route
	routes = append(routes, LoginRoutes...)
	routes = append(routes, UserRoutes...)
	routes = append(routes, PublicationsRoutes...)
	routes = append(routes, HomeRoute...)

	for _, route := range routes {
		if route.Authentication {
			router.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))
}

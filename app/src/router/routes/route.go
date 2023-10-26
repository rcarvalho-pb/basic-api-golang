package routes

import (
	"net/http"

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

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))
}

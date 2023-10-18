package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	var routes []Route
	routes = append(routes, UserRoutes...)
	routes = append(routes, AuthenticationRoutes...)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}

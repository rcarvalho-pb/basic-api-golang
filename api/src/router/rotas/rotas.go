package rotas

import (
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
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	router := mux.NewRouter()
	routes.Config(router)
	
	return router
}
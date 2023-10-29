package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var HomeRoute = Route{
	Uri: "/home",
	Method: http.MethodGet,
	Function: controllers.LoadHome,
	Authentication: true,
}
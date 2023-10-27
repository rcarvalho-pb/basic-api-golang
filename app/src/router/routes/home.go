package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var HomeRoute = Route{
	Uri: "/huome",
	Method: http.MethodGet,
	Function: controllers.LoadHome,
	Authentication: false,
}
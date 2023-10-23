package routes

import (
	"net/http"
	"webapp/src/controllers"
)

const LOGIN_RESOURCE = "/login"

var LoginRoutes = []Route {
	{
		Uri: "/",
		Method: http.MethodGet,
		Function: controllers.Login,
		Authentication: false,
	},
	{
		Uri: LOGIN_RESOURCE,
		Method: http.MethodGet,
		Function: controllers.Login,
		Authentication: false,
	},
}
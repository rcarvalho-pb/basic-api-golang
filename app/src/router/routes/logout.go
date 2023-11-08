package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var LogoutRoutes = []Route{
	{
		Uri: "/logout",
		Method: http.MethodGet,
		Function: controllers.Logout,
		Authentication: true,
	},
}
package rotas

import (
	"api/src/controllers"
	"net/http"
)

const BASE_URL = "/auth"

var AuthenticationRoutes = []Route{
	{
		Uri: "/login",
		Method: http.MethodPost,
		Function: controllers.Login,
		RequireAuthentication: false,
	},
}
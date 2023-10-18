package rotas

import (
	"api/src/controllers"
	"net/http"
)

const AUTH_RESOURCE = "/auth"

var AuthenticationRoutes = []Route{
	{
		Uri:                   AUTH_RESOURCE + "/login",
		Method:                http.MethodPost,
		Function:              controllers.Login,
		RequireAuthentication: false,
	},
}

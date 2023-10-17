package rotas

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.GetUserById,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUserById,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUserById,
		RequireAuthentication: false,
	},
}

package rotas

import (
	"api/src/controllers"
	"net/http"
)

const USER_RESOURCE = "/users"

var UserRoutes = []Route{
	{
		Uri:                   USER_RESOURCE,
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   USER_RESOURCE,
		Method:                http.MethodGet,
		Function:              controllers.FindUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.GetUserById,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUserById,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUserById,
		RequireAuthentication: true,
	},
}

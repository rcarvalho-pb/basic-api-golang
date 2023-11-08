package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var UserRoutes = []Route {
	{
		Uri: "/create-user",
		Method: http.MethodGet,
		Function: controllers.LoadUserRegisterPage,
		Authentication: false,
	},
	{
		Uri: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		Authentication: false,
	},
	{
		Uri: "/search-user",
		Method: http.MethodGet,
		Function: controllers.LoadUsersPage,
		Authentication: true,
	},
	{
		Uri: "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.LoadUserProfile,
		Authentication: true,
	},
}
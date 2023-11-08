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
	{
		Uri:                   USER_RESOURCE + "/{userId}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnfollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}/followers",
		Method:                http.MethodGet,
		Function:              controllers.UserFollowers,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}/followed",
		Method:                http.MethodGet,
		Function:              controllers.UserFollowed,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}

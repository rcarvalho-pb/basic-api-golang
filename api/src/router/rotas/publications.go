package rotas

import (
	"api/src/controllers"
	"net/http"
)

const PUBLICATIOIN_RESOURCE = "/publications"

var PublicationsRoutes = []Route{
	{
		Uri:                   PUBLICATIOIN_RESOURCE,
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   PUBLICATIOIN_RESOURCE,
		Method:                http.MethodGet,
		Function:              controllers.FindPublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   PUBLICATIOIN_RESOURCE + "/{publicationId}",
		Method:                http.MethodGet,
		Function:              controllers.FindPublicationById,
		RequireAuthentication: true,
	},
	{
		Uri:                   PUBLICATIOIN_RESOURCE + "/{publicationId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   PUBLICATIOIN_RESOURCE + "/{publicationId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   USER_RESOURCE + "/{userId}" + PUBLICATIOIN_RESOURCE,
		Method:                http.MethodGet,
		Function:              controllers.GetUserPublications,
		RequireAuthentication: true,
	},
	{
		Uri:                   PUBLICATIOIN_RESOURCE + "/{publicationId}/like",
		Method:                http.MethodPatch,
		Function:              controllers.LikePublication,
		RequireAuthentication: true,
	},
	{
		Uri:                   PUBLICATIOIN_RESOURCE + "/{publicationId}/dislike",
		Method:                http.MethodPatch,
		Function:              controllers.DislikePublication,
		RequireAuthentication: true,
	},
}
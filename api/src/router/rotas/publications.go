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
}
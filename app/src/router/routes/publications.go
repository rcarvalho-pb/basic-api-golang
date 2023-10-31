package routes

import (
	"net/http"
	"webapp/src/controllers"
)

const PUBLICATIONS_RESOURCES = "/publications"

var PublicationsRoutes = []Route {
	{
		Uri: PUBLICATIONS_RESOURCES,
		Method: http.MethodPost,
		Function: controllers.CreatePublication,
		Authentication: true,
	},
	{
		Uri: PUBLICATIONS_RESOURCES + "/{publicationId}/like",
		Method: http.MethodPatch,
		Function: controllers.LikePublication,
		Authentication: true,
	},
	{
		Uri: PUBLICATIONS_RESOURCES + "/{publicationId}/dislike",
		Method: http.MethodPatch,
		Function: controllers.DislikePublication,
		Authentication: true,
	},
	{
		Uri: PUBLICATIONS_RESOURCES + "/{publicationId}/update-publication",
		Method: http.MethodGet,
		Function: controllers.LoadUpdatePublicationPage,
		Authentication: true,
	},
	{
		Uri: PUBLICATIONS_RESOURCES + "/{publicationId}",
		Method: http.MethodPut,
		Function: controllers.UpdatePublication,
		Authentication: true,
	},
	{
		Uri: PUBLICATIONS_RESOURCES + "/{publicationId}",
		Method: http.MethodDelete,
		Function: controllers.DeletetePublication,
		Authentication: true,
	},
}
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
}
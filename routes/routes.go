package routes

import (
	"net/http"

	"github.com/enricod/1h1dphoto.com-be/html"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/enricod/1h1dphoto.com-be/rest"
	"github.com/gorilla/mux"
)

// Route definizione di una route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes elenco delle routes
type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"status",
		"GET",
		model.API_PREFIX + "/status",
		rest.Status,
	},
	Route{
		"userRegister",
		"POST",
		model.API_PREFIX + "/users/register",
		rest.UserRegister,
	},
	Route{
		"userCodeValidation",
		"POST",
		model.API_PREFIX + "/users/codeValidation",
		rest.UserCodeValidation,
	},
	Route{
		"eventsSummaryList",
		"GET",
		model.API_PREFIX + "/events/summary/list",
		rest.IsAuthenticated(rest.EventsSummary),
	},
	Route{
		"event",
		"GET",
		model.API_PREFIX + "/events/{eventID}",
		rest.IsAuthenticated(rest.Event),
	},
	Route{
		"logout",
		"GET",
		model.API_PREFIX + "/users/logout/{token}",
		rest.Logout,
	},
	Route{
		"imgUpload",
		"POST",
		model.API_PREFIX + "/images/upload",
		rest.IsAuthenticated(rest.ImgUpload),
	},
	Route{
		"imgDownload",
		"GET",
		model.API_PREFIX + "/images/download/{id}",
		rest.ImgDownload,
	},
	Route{
		"home",
		"GET",
		"/",
		html.Home,
	},
}

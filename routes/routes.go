package routes

import (
	"net/http"

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
		"userRegister",
		"POST",
		"/api/users/register",
		rest.UserRegister,
	},
	Route{
		"userCodeValidation",
		"POST",
		"/api/users/codeValidation",
		rest.UserCodeValidation,
	},
	Route{
		"eventsSummaryList",
		"GET",
		"/api/events/summary/list",
		rest.IsAuthenticated(rest.EventsSummary),
	},
	Route{
		"event",
		"GET",
		"/api/events/{eventID}",
		rest.IsAuthenticated(rest.Event),
	},
	Route{
		"logout",
		"GET",
		"/api/users/logout/{token}",
		rest.Logout,
	},
	Route{
		"imgUpload",
		"POST",
		"/api/images/upload",
		rest.IsAuthenticated(rest.ImgUpload),
	},
	Route{
		"imgDownload",
		"GET",
		"/api/images/download/{id}",
		rest.ImgDownload,
	},
}

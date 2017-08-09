package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/enricod/1h1dphoto.com-be/rest"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
		"/user/register",
		rest.UserRegister,
	},
	Route{
		"userCodeValidation",
		"POST",
		"/user/codeValidation",
		rest.UserCodeValidation,
	},
	Route{
		"eventsList",
		"GET",
		"/events/list",
		rest.IsAuthenticated(rest.Events),
	},
	Route{
		"logout",
		"GET",
		"/logout/{token}",
		rest.Logout,
	},
	Route{
		"sessionList",
		"GET",
		"/sessions",
		rest.Sessions,
	},
}


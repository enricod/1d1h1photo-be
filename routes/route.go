package routes

import (
	"net/http"

	"1d1hphoto.com-be/rest"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"login",
		"GET",
		"/login/{username}/{password}",
		rest.LoginReq,
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

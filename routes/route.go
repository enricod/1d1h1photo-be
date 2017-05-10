package routes

import (
	"../rest"
	"net/http"
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
		"/logout",
		rest.LoginReq,
	},
}

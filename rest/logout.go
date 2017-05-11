package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Logout(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	Tokens[username] = nil
}

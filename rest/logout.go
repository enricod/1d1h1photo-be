package rest

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"github.com/enricod/1h1dphoto.com-be/model"
)

func Logout(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	Tokens[username] = nil
}

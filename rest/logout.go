package rest

import (
<<<<<<< HEAD:rest/logout.go
	"fmt"
=======
>>>>>>> c4cfc481525c81a69bd4c9082683268d887e88d0:src/it/1d1hphoto/rest/logout.go
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

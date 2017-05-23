package rest

import (
	"net/http"

	"1d1hphoto.com-be/model"
	"1d1hphoto.com-be/service"

	"fmt"

	"github.com/gorilla/mux"
)

func Registra(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	res.Header().Set("Content-Type", "application/json")
	user := model.User{}
	user.Name = username
	id := service.InsertUser(user)
	fmt.Println("inserito id ", id)
}

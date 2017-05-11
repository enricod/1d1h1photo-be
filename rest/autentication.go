package rest

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"1d1hphoto.com-be/model"
	"1d1hphoto.com-be/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var Tokens = make(map[string]model.User)

func LoginReq(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	password := vars["password"]
	res.Header().Set("Content-Type", "application/json")
	user := autentica(username, password)
	if user.Id == 0 {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := createToken(username, req.Host)
	sToken, _ := token.SignedString([]byte("secret"))
	Tokens[sToken] = user
	res.WriteHeader(http.StatusOK)
	err := json.NewEncoder(res).Encode(model.Response{Data: sToken})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

}

func Logout(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sToken := vars["token"]
	delete(Tokens, sToken)
	res.WriteHeader(http.StatusOK)
}

func Sessions(res http.ResponseWriter, req *http.Request) {
	s := reflect.ValueOf(Tokens).MapKeys()
	var str string
	for _, element := range s {
		str += element.String() + ","
	}
	err := json.NewEncoder(res).Encode(model.Response{Data: str})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func createToken(username string, host string) *jwt.Token {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claim := model.Claims{
		username,
		time.Now().Round(time.Millisecond).UnixNano(),
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    host,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token
}

func autentica(username string, password string) model.User {
	user, err := service.GetUser(username)
	if err != nil {
		return model.User{}
	}
	if user.Password == password {
		return user
	}
	return model.User{}
}

package rest

import (
	"encoding/json"
	"fmt"
	"it/1d1hphoto/model"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var Tokens = make(map[string]*jwt.Token)

func LoginReq(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	password := vars["password"]
	fmt.Println(password)
	res.Header().Set("Content-Type", "application/json")

	sToken := ""
	if Tokens[username] == nil {
		if autentica(username, password) == false {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("create token")
		sToken = createToken(username, req.Host)
	} else {
		sToken, _ = Tokens[username].SignedString([]byte("secret"))
	}

	res.WriteHeader(http.StatusOK)
	err := json.NewEncoder(res).Encode(model.Response{Data: sToken})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

}

func createToken(username string, host string) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	claim := model.Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    host,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	Tokens[username] = token

	signedToken, _ := token.SignedString([]byte("secret"))

	return signedToken
}

func autentica(username string, password string) bool {
	return false
}

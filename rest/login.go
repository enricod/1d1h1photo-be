package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/enricod/1h1dphoto.com-be/model"
)

func LoginReq(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	password := vars["password"]
	fmt.Println(password)
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	claim := model.Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    req.Host,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, _ := token.SignedString([]byte("secret"))

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(model.Response{Data: signedToken})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

}

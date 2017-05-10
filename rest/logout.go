package rest

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"time"
	"github.com/enricod/1h1dphoto.com-be/model"
)

func Logout(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	signedToken := vars["token"]
	token, err := jwt.ParseWithClaims(signedToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte("secret"), nil
	})
	token.Claims.Set(ExpiresAt:time.Now())
	res.Header().Set("Content-Type", "application/json")
}

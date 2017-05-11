package rest

import (
	"context"
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/enricod/1h1dphoto.com-be/model"
)

func Validate(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		signedToken := vars["Auth"]

		token, err := jwt.ParseWithClaims(signedToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claim, ok := token.Claims.(*model.Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), model.MyKey, *claim)
			protectedPage(res, req.WithContext(ctx))
		} else {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}

//claims, ok := req.Context().Value(MyKey).(Claims)

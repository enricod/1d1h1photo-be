package rest

import (
	"net/http"
	"github.com/enricod/1h1dphoto.com-be/db"

)

func IsAuthenticated(f func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		appAuth := req.Header.Get("Auth")

		userAppToken := db.FindAppToken(appAuth)
		if userAppToken == nil || !userAppToken.Valid {
			w.WriteHeader(401)
		}

		f(w, req)
	}
}

func Events(res http.ResponseWriter, req *http.Request) {



}
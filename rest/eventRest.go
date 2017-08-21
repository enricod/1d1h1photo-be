package rest

import (
	"net/http"
	"github.com/enricod/1h1dphoto.com-be/db"
	"encoding/json"
	"github.com/enricod/1h1dphoto.com-be/model"
)

func IsAuthenticated(f func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		appAuth := req.Header.Get("Authorization")

		userAppToken := db.FindAppToken(appAuth)
		if userAppToken == nil || !userAppToken.Valid {
			w.WriteHeader(401)
		} else {
			f(w, req)
		}
	}
}

func EventsSummary(res http.ResponseWriter, req *http.Request) {

	events := db.EventsList(4)

	// separiamo eventi in 2 blocchi: quello futuro, e l'elenco di quelli passati
	var closed = make( []db.Event, 0)
	var next db.Event
	for _, e := range events {
		if e.IsClosed() {
			closed = append(closed, e)
		} else {
			next = e
		}
	}

	body := model.EventsSummaryResBody{ClosedEvents: closed, NextEvent:next}
	eventsListRes := model.EventsListRes{Body:body}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(res).Encode(eventsListRes)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

}
package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enricod/1h1dphoto.com-be/db"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/gorilla/mux"
)

//
// Confs configurazioni applicazione
//
var Confs model.AppConfs

//
// IsAuthenticated utilizzata per controllare se token applicazione è valido
//
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

func Event(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	eventId, _ := strconv.Atoi(vars["eventId"])
	event, _ := db.EventDetails(uint(eventId))

	response := model.EventResBody{Body: event}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(res).Encode(response)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func EventsSummary(res http.ResponseWriter, req *http.Request) {

	events := db.EventsList(4)

	// separiamo eventi in 2 blocchi: quello futuro, e l'elenco di quelli passati
	var closed = make([]model.Event, 0)
	var next model.Event
	for _, e := range events {
		if e.IsClosed() {
			closed = append(closed, e)
		} else {
			next = e
		}
	}

	body := model.EventsSummaryResBody{ClosedEvents: closed, NextEvent: next}
	eventsListRes := model.EventsListRes{Body: body}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(res).Encode(eventsListRes)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

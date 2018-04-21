package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enricod/1h1dphoto.com-be/db"
	"github.com/enricod/1h1dphoto.com-be/model"
	"github.com/gorilla/mux"
)

// IsAuthenticated utilizzata per controllare se token applicazione è valido. Mi aspetto un header Authorization contentente appToken.
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

// Event dettaglio un evento con immagini collegate
func Event(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	eventID, _ := strconv.Atoi(vars["eventID"])
	event, _ := db.EventDetails(uint(eventID))

	response := model.EventResBody{Body: event}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(res).Encode(response)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

// EventsSummary elenco eventi con 3 immagini per ogni evento
func EventsSummary(res http.ResponseWriter, req *http.Request) {
	var nrImmaginiPerEvento = 3
	events := db.EventsList(nrImmaginiPerEvento)
	// separiamo eventi in 2 blocchi: quello futuro, e l'elenco di quelli passati
	var closed = make([]model.Event, 0)
	var future = make([]model.Event, 0)
	for _, e := range events {
		if e.IsClosed() {
			closed = append(closed, e)
		} else {
			future = append(future, e)
		}
	}

	body := model.EventsSummaryResBody{ClosedEvents: closed, FutureEvents: future}
	eventsListRes := model.EventsListRes{Body: body}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(res).Encode(eventsListRes)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

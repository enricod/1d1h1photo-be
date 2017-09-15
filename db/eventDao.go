package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/enricod/1h1dphoto.com-be/model"
)

// FindEventSubmissions elenco delle immagini associate a un evento
func FindEventSubmissions(event *model.Event, limit int) []model.EventSubmission {
	db := openDB()
	var eventSubmissions []model.EventSubmission
	db.Model(&event).Limit(limit).Order("score", true).Related(&eventSubmissions)
	return eventSubmissions
}

// EventsList elenco degli eventi.
func EventsList(limit int) []model.Event {
	db := openDB()
	var events []model.Event
	db.Order("start DESC").Find(&events)
	for k := range events {
		submissions := FindEventSubmissions(&events[k], limit)
		for _, s := range submissions {
			s.ThumbUrl = "/images/download/" + fmt.Sprint(s.ID) + "?size=t"
			s.ImageUrl = "/images/download/" + fmt.Sprint(s.ID) + "?size=s"
			events[k].Submissions = append(events[k].Submissions, s)
		}
	}
	return events
}

// FindEventoByDate cerca app token
func FindEventoByDate(data time.Time) *model.Event {
	db := openDB()
	var result model.Event
	db.First(&result, " start <= ? AND end > ? ", data, data)
	if result.ID > 0 {
		return &result
	}
	return nil
}

// SubmissionByID ritorna una foto dato l'ID
func SubmissionByID(id uint) (model.EventSubmission, error) {
	db := openDB()
	var submission model.EventSubmission
	db.First(&submission, id) // find product with code l1212
	if submission.ID > 0 {
		return submission, nil
	}
	return submission, errors.New("nessun utente con email ")
}

// EventDetails dettaglio evento (carica 10000 immagini)
func EventDetails(eventID uint) (model.Event, error) {
	db := openDB()
	var event model.Event
	db.First(&event, eventID)
	if event.ID > 0 {
		submissions := FindEventSubmissions(&event, 10000)
		//fmt.Println("trovati %v submissions", len(submissions))
		for _, s := range submissions {
			s.ThumbUrl = "/images/download/" + fmt.Sprint(s.ID) + "?size=t"
			s.ImageUrl = "/images/download/" + fmt.Sprint(s.ID) + "?size=s"
			event.Submissions = append(event.Submissions, s)
		}
		return event, nil
	}
	return event, errors.New("nessun utente con email ")
}

// InsertSubmission salva nel DB un'immagine
func InsertSubmission(eventID uint, token *model.UserAppToken, imageUID string, imageName string) bool {
	db := openDB()

	eventSubmission := model.EventSubmission{
		EventId:        eventID,
		UserId:         token.UserId,
		ImageName:      imageName,
		ImageUid:       imageUID,
		SubmissionDate: time.Now()}

	db.Create(&eventSubmission)

	return true

}

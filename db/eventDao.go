package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/enricod/1h1dphoto.com-be/model"
)

func FindEventSubmissions(event *model.Event, limit int) []model.EventSubmission {
	db := openDB()
	var eventSubmissions []model.EventSubmission
	db.Model(&event).Limit(limit).Order("score", true).Related(&eventSubmissions)
	return eventSubmissions
}

func EventsList(limit uint) []model.Event {
	db := openDB()
	var events []model.Event
	db.Order("start DESC").Find(&events)
	for k, _ := range events {
		submissions := FindEventSubmissions(&events[k], 3)
		//fmt.Println("trovati submissions: ", len(submissions))
		for _, s := range submissions {
			s.ThumbUrl = "/api/images/download/" + fmt.Sprint(s.ID) + "?size=t"
			s.ImageUrl = "/api/images/download/" + fmt.Sprint(s.ID) + "?size=s"
			events[k].Submissions = append(events[k].Submissions, s)
		}
	}
	return events
}

func SubmissionById(id uint) (model.EventSubmission, error) {
	db := openDB()
	var submission model.EventSubmission
	db.First(&submission, id) // find product with code l1212
	if submission.ID > 0 {
		return submission, nil
	} else {
		return submission, errors.New("nessun utente con email ")
	}
}

func EventDetails(eventId uint) (model.Event, error) {
	db := openDB()
	var event model.Event
	db.First(&event, eventId)
	if event.ID > 0 {
		submissions := FindEventSubmissions(&event, 10000)
		//fmt.Println("trovati %v submissions", len(submissions))
		for _, s := range submissions {
			s.ThumbUrl = "/api/images/download/" + fmt.Sprint(s.ID) + "?size=t"
			s.ImageUrl = "/api/images/download/" + fmt.Sprint(s.ID) + "?size=s"
			event.Submissions = append(event.Submissions, s)
		}
		return event, nil
	} else {
		return event, errors.New("nessun utente con email ")
	}
}

func InsertSubmission(eventId uint, token *model.UserAppToken, imageUid string, imageName string) bool {
	db := openDB()

	eventSubmission := model.EventSubmission{
		EventId:        eventId,
		UserId:         token.UserId,
		ImageName:      imageName,
		ImageUid:       imageUid,
		SubmissionDate: time.Now()}

	db.Create(&eventSubmission)

	return true

}

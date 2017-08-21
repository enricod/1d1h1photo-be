package db

import (
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)


type Event struct {
	gorm.Model
	Start time.Time
	End time.Time
	Name string
	Submissions []EventSubmission
}

func ( e Event) IsClosed() bool {
	return e.End.Before( time.Now() )
}



type EventSubmission struct {
	gorm.Model
	Event Event `json:"-"`
	EventId uint
	User User `json:"-"`
	UserId uint
	UploadTime time.Time
	ImageName string
	LikesNr uint
	Score float32
	Latitude float32
	Longitude float32
	SubmissionDate time.Time
}

type EventSubmissionAction struct {
	gorm.Model
	User User `gorm:"ForeignKey:UserId"`
	UserId uint
	Time time.Time
	EventSubmission EventSubmission `gorm:"ForeignKey:UserId"`
	EventSubmissionId uint
	Type string
}

func FindEventSubmissions( event *Event) []EventSubmission {
	db := openDB();
	var eventSubmissions []EventSubmission
	db.Model(&event).Related(&eventSubmissions)
	return eventSubmissions
}

func EventsList(limit uint) []Event {
	db := openDB();
	var events []Event
	db.Limit(limit).Order("start DESC").Find(&events)
	for k, _ := range events {
		submissions :=  FindEventSubmissions(&events[k])
		fmt.Println("trovati %v submissions", len(submissions))
		for _, s := range submissions {
			 events[k].Submissions = append(events[k].Submissions, s)
		}
	}

	return events
}



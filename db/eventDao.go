package db

import (
	"github.com/jinzhu/gorm"
	"time"
)


type Event struct {
	gorm.Model
	Start time.Time
	End time.Time
	Name string
}

type EventSubmission struct {
	gorm.Model
	Event Event `gorm:"ForeignKey:EventId"`
	EventId uint
	User User `gorm:"ForeignKey:UserId"`
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




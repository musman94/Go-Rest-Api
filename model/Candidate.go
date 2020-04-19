package model

import (
	"time"
)

//Candidate Model
type Candidate struct {
	ID              string     `bson:"_id" json:"_id"`
	FirstName       string     `bson:"first_name" json:"first_name"`
	LastName        string     `bson:"last_name" json:"last_name"`
	Email           string     `bson:"email" json:"email"`
	Department      string     `bson:"department" json:"department"`
	University      string     `bson:"university" json:"university"`
	Experience      bool       `bson:"experience" json:"experience"`
	Status          string     `bson:"status" json:"status"`
	MeetingCount    int        `bson:"meeting_count" json:"meeting_count"`
	NextMeeting     *time.Time `bson:"next_meeting" json:"next_meeting"`
	Assignee        string     `bson:"assignee" json:"assignee"`
	ApplicationDate string     `bson:"application_date" json:"application_date"`
}

package service

import (
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	. "github.com/musman94/Otsimo-Task/helper"
	. "github.com/musman94/Otsimo-Task/model"
	"github.com/musman94/Otsimo-Task/repository"
)

//Departments to which the candidate can apply to
var departments = map[string]bool{"Marketing": true, "Design": true, "Development": true}

//CreateCandidate service layer function
//Validates the input and then passes the candidate to the repository layer function
func CreateCandidate(candidate Candidate) (Candidate, error) {
	if err := validateInput(candidate); err != nil {
		return candidate, err
	}

	candidate.ID = bson.NewObjectId().Hex()
	candidate.MeetingCount = 0
	candidate.Status = "Pending"
	candidate.NextMeeting = nil
	candidate.ApplicationDate = time.Now().Format("02-01-2006")

	result, err := repository.CreateCandidate(candidate)

	if err != nil {
		return result, &HttpError{http.StatusInternalServerError, "There was an internal problem while creating the candidate"}
	}

	return result, err
}

//ReadCandidate service layer function
//Gets the candidate by passing the ID to the repository layer function
func ReadCandidate(_id string) (Candidate, error) {
	var result Candidate
	var err error

	if result, err = repository.ReadCandidate(_id); err != nil {
		return result, &HttpError{http.StatusNotFound, "Invalid candidate ID, candidate couldn't be found"}
	}

	return result, err
}

func ReadAllCandidates() (*mongo.Cursor, error) {
	var cur *mongo.Cursor
	var err error

	if cur, err = repository.ReadAllCandidates(); err != nil {
		return cur, &HttpError{http.StatusInternalServerError, "There was an internal problem while fetching the assignee's candidates"}
	}

	return cur, err
}

//DeleteCandidate service layer function
//Deletes the candidate by passing the ID to the repository layer function
func DeleteCandidate(_id string) error {
	if _, err := ReadCandidate(_id); err != nil {
		return &HttpError{http.StatusNotFound, "Invalid candidate ID, candidate couldn't be found"}
	}

	if err := repository.DeleteCandidate(_id); err != nil {
		return &HttpError{http.StatusInternalServerError, "There was an internal problem while deleting the candidate"}
	}

	return nil
}

//ArrangeMeeting service layer function
//Fetches the candidate by calling the ReadCandidate service layer function
//Changes the candidate's fields
//Passes the candidate to the repository layer UpdateCandidate function
func ArrangeMeeting(_id string, nextMeetingTime *time.Time) error {
	var candidate Candidate
	var err error

	if candidate, err = ReadCandidate(_id); err != nil {
		return &HttpError{http.StatusNotFound, "Invalid candidate ID, candidate couldn't be found"}
	}

	if candidate.MeetingCount == 4 {
		return &HttpError{http.StatusBadRequest, "More than 4 meetings with the candidate are not allowed"}
	}

	if candidate.Status == "Pending" {
		candidate.Status = "In Progress"
	}

	candidate.NextMeeting = nextMeetingTime

	if err = repository.UpdateCandidate(candidate); err != nil {
		return &HttpError{http.StatusInternalServerError, "There was an internal problem while arranging a meeting with the candidate"}
	}

	return nil
}

//CompleteMeeting service layer function
//Fetches the candidate by calling the ReadCandidate service layer function
//Changes the candidate's fields
//Passes the candidate to the repository layer UpdateCandidate function
func CompleteMeeting(_id string) error {
	var candidate Candidate
	var err error

	if candidate, err = ReadCandidate(_id); err != nil {
		return &HttpError{http.StatusNotFound, "Invalid candidate ID, candidate couldn't be found"}
	}

	if candidate.MeetingCount == 4 {
		return &HttpError{http.StatusBadRequest, "The candidate has already completed all of the meetings"}
	}

	if candidate.NextMeeting == nil {
		return &HttpError{http.StatusBadRequest, "The candidate has no meetings in progress"}
	}

	candidate.MeetingCount++
	candidate.NextMeeting = nil

	if candidate.MeetingCount == 3 {
		candidate.Assignee = FindAssigneeIDByName("Zafer")
	}

	if err = repository.UpdateCandidate(candidate); err != nil {
		return &HttpError{http.StatusInternalServerError, "There was an internal problem while completing the meeting with the candidate"}
	}

	return nil
}

//DenyCandidate service layer function
//Fetches the candidate by calling the ReadCandidate service layer function
//Changes the candidate's fields
//Passes the candidate to the repository layer UpdateCandidate function
func DenyCandidate(_id string) error {
	var candidate Candidate
	var err error

	if candidate, err = ReadCandidate(_id); err != nil {
		return &HttpError{http.StatusNotFound, "Invalid candidate ID, candidate couldn't be found"}
	}

	candidate.Status = "Denied"

	if err = repository.UpdateCandidate(candidate); err != nil {
		return &HttpError{http.StatusInternalServerError, "There was an internal problem while completing the meeting with the candidate"}
	}

	return nil
}

//AcceptCandidate service layer function
//Fetches the candidate by calling the ReadCandidate service layer function
//Changes the candidate's fields
//Passes the candidate to the repository layer UpdateCandidate function
func AcceptCandidate(_id string) error {
	var candidate Candidate
	var err error

	if candidate, err = ReadCandidate(_id); err != nil {
		return &HttpError{http.StatusNotFound, "Invalid candidate ID, candidate couldn't be found"}
	}

	if candidate.MeetingCount < 4 {
		return &HttpError{http.StatusBadRequest, "Candidate cannot be accepted before having 4 meetings"}
	}

	candidate.Status = "Accepted"

	if err = repository.UpdateCandidate(candidate); err != nil {
		return &HttpError{http.StatusInternalServerError, "There was an internal problem while completing the meeting with the candidate"}
	}

	return nil
}

//Checks if email has the right format
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

//Validates the candidate's body
//Returns an error depending on the problem
func validateInput(candidate Candidate) error {
	if !validateEmail(candidate.Email) {
		return &HttpError{http.StatusBadRequest, "Invalid email"}
	}

	if !departments[candidate.Department] {
		return &HttpError{http.StatusBadRequest, "The department applied to does not exist"}
	}

	assignee, err := FindAssigneeByID(candidate.Assignee)

	if err != nil {
		return &HttpError{http.StatusBadRequest, "Assignee not found"}
	}

	if candidate.Department != assignee.Department {
		return &HttpError{http.StatusBadRequest, "Assignee should belong to the department the candidate applied to"}
	}

	return nil
}

package service

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/musman94/Otsimo-Task/helper"
	. "github.com/musman94/Otsimo-Task/model"
	"github.com/musman94/Otsimo-Task/repository"
)

/*
func CreateAssignee(assignee Assignee) (Assignee, error) {
	assignee.ID = primitive.NewObjectID().String()

	result, err := repository.CreateAssignee(assignee)

	return result, err
}
*/

//FindAssigneeByID service layer function
//Fetches the assignee using its ID by calling the repository layer function
func FindAssigneeByID(_id string) (Assignee, error) {
	var assignee Assignee
	var err error

	if assignee, err = repository.ReadAssigneeByID(_id); err != nil {
		return assignee, &HttpError{http.StatusNotFound, "Invalid assignee ID, assignee couldn't be found"}
	}

	return assignee, err
}

//FindAssigneeIDByName service layer function
//Fetches the assignee's ID using its name by calling the repository layer function
func FindAssigneeIDByName(name string) string {
	var assignee Assignee
	var err error

	if assignee, err = repository.ReadAssigneeByName(name); err != nil {
		return ""
	}

	return assignee.ID
}

//FindAssigneesCandidates service layer function
//Fetches all of the candidates by calling the ReadAllCandidates repository layer function
//Adds the candidates to the result which have the assignee ID of _id
func FindAssigneesCandidates(_id string) ([]Candidate, error) {
	var candidates []Candidate
	var cur *mongo.Cursor
	var err error

	if _, err = FindAssigneeByID(_id); err != nil {
		return candidates, &HttpError{http.StatusNotFound, "Invalid assignee ID, assignee couldn't be found"}
	}

	if cur, err = repository.ReadAllCandidates(); err != nil {
		return candidates, err
	}

	for cur.Next(context.TODO()) {
		var candidate Candidate
		err := cur.Decode(&candidate)

		if err != nil {
			return candidates, err
		}

		if candidate.Assignee == _id {
			candidates = append(candidates, candidate)
		}
	}

	cur.Close(context.TODO())

	return candidates, err
}

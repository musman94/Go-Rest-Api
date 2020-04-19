package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	. "github.com/musman94/Otsimo-Task/helper"
	. "github.com/musman94/Otsimo-Task/model"
)

var Candidates *mongo.Collection

//CreateCandidate dao layer function
//Creates and returns the candidate from the database
func CreateCandidate(candidate Candidate) (Candidate, error) {
	_, err := Candidates.InsertOne(context.TODO(), candidate)

	return candidate, err
}

//ReadCandidate dao layer function
//Fetches and returns the candidate from the database
func ReadCandidate(_id string) (Candidate, error) {
	var candidate Candidate

	filter := bson.M{"_id": _id}

	err := Candidates.FindOne(context.TODO(), filter).Decode(&candidate)

	return candidate, err
}

//ReadAllCandidates dao layer function
//Fetches and returns all of the candidates from the database
func ReadAllCandidates() (*mongo.Cursor, error) {
	filter := bson.M{}

	cur, err := Candidates.Find(context.TODO(), filter)

	return cur, err
}

//UpdateCandidate dao layer function
//Updates the candidate using it's ID
func UpdateCandidate(candidate Candidate) error {
	filter := bson.M{"_id": candidate.ID}

	updated := bson.M{
		"$set": candidate,
	}

	_, err := Candidates.UpdateOne(context.TODO(), filter, updated)

	return err
}

//DeleteCandidate dao layer function
//Deletes the candidate using it's ID
func DeleteCandidate(_id string) error {
	filter := bson.M{"_id": _id}

	_, err := Candidates.DeleteOne(context.TODO(), filter)

	return err
}

//Connects to the database and gets the Candidates collection
func init() {
	Candidates = ConnectDB("Candidates")
}

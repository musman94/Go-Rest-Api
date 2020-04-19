package repository

import (
	"context"

	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/musman94/Otsimo-Task/helper"
	. "github.com/musman94/Otsimo-Task/model"
)

var Assignees *mongo.Collection

//ReadAssigneeByID dao layer function
//Fetches and returns the assignee from the database
func ReadAssigneeByID(_id string) (Assignee, error) {
	var assignee Assignee

	filter := bson.M{"_id": _id}

	err := Assignees.FindOne(context.TODO(), filter).Decode(&assignee)

	return assignee, err
}

//ReadAssigneeByName dao layer function
//Fetches and returns the assignee's ID from the database
func ReadAssigneeByName(name string) (Assignee, error) {
	var assignee Assignee

	filter := bson.M{"name": name}

	err := Assignees.FindOne(context.TODO(), filter).Decode(&assignee)

	return assignee, err
}

//Connects to the database and gets the Assignees collection
func init() {
	Assignees = ConnectDB("Assignees")
}

package helper

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectDB helper function to connect to the database and get the required collection using the parameter "collectionName"
func ConnectDB(collectionName string) *mongo.Collection {
	//Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//Fetch the collection
	collection := client.Database("Otsimo").Collection(collectionName)

	return collection
}

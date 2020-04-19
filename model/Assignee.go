package model

//Assignee Model
type Assignee struct {
	ID         string `bson:"_id" json:"_id"`
	Name       string `bson:"name" json:"name"`
	Department string `bson:"department" json:"department"`
}

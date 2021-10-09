package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` 
	UserID 			string 		`json:"userid,omitempty" bson:"userid,omitempty"`
	Caption         string		`json:"Caption,omitempty" bson:"Caption,omitempty"`
	ImageURL        string		`json:"ImageURL,omitempty" bson:"ImageURL,omitempty"`
}
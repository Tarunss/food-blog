package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Summary string             `json:"summary,omitempty"`
	Body    string             `json:"body,omitempty"`
	Date    string             `json:"date,omitempty"`
	//Image []byte             `json:"image,omitempty"`
}

type UsePass struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

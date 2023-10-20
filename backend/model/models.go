package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Body string             `json:"body,omitempty"`
	Date string             `json:"date,omitempty"`
}

package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

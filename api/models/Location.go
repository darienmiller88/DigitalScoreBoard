package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"` 
	LocationName string             `bson:"location_name" json:"location_name"`
	Users        []string           `bson:"users"         json:"users"`
}
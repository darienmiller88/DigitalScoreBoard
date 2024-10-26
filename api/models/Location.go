package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"` 
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	LocationName string             `bson:"location_name" json:"location_name"`
	Users        []UserCard         `bson:"users"         json:"users"`
}
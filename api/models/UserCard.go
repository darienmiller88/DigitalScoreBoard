package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCard struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"name"`
	Score     int                `bson:"score,omitempty"`
}
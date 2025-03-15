package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-ozzo/ozzo-validation"
)

type Location struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"` 
	CreatedAt    time.Time          `bson:"created_at"    json:"created_at"` 
	UpdatedAt    time.Time          `bson:"updated_at"    json:"updated_at"`
	LocationName string             `bson:"location_name" json:"location_name"`
	Users        []UserCard         `bson:"users"         json:"users"`
}

func (l *Location) InitCreatedAtAndUpdatedAt(){
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()

	//If this field is not initialized, it is interpreted as "null" by mongoDB, and not an empty array.
	l.Users = []UserCard{}
}

func (l *Location) Validate() error{
	return validation.ValidateStruct(
		l,
		validation.Field(&l.LocationName, validation.Required),
	)
}

// func (l *Location) findLocation() error{

// }
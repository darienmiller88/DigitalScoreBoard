package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-ozzo/ozzo-validation"
)

type SavedGame struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // The MongoDB document ID
	Location Location           `bson:"location,inline" json:"location"`
	Winner   string             `bson:"winner"          json:"winner"`
}

func (s *SavedGame) Validate() error{
	return validation.ValidateStruct(
		s,
		validation.Field(s.Winner, validation.By(s.findWinner)),
	)
}

func ( *SavedGame) findWinner(s interface{}) error{
	winner, ok := s.(string)

	if !ok{
		return fmt.Errorf("could not parse %s into string", s)
	}

	fmt.Println(winner)

	return nil
}
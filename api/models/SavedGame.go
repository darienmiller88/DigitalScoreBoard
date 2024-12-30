package models

import (
	"fmt"
	"time"
	"slices"

	"github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SavedGame struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"` // The MongoDB document ID
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	Location      Location           `bson:",inline"        json:"location"`
	TotalPoints   int 			     `bson:"total_points"   json:"total_points"`
	AveragePoints float64            `bson:"average_points" json:"average_points"`
	Winner        UserCard           `bson:"winner"         json:"winner"`
}

func (s *SavedGame) Validate() error{
	return validation.ValidateStruct(
		s,
		validation.Field(s.Winner, validation.By(s.findWinner)),
	)
}

func (s *SavedGame) InitCreatedAtAndUpdatedAt(){
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	s.Location.InitCreatedAtAndUpdatedAt()
}

func (s *SavedGame) findWinner(field interface{}) error{
	winner, ok := field.(UserCard)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}

	if !slices.Contains(s.Location.Users, winner){
		return fmt.Errorf("could not find winner %s in list of players", winner.Name)
	}

	return nil
}

func (s *SavedGame) CalcAveragePoints(){
	if len(s.Location.Users) == 0 {
		return
	}

	if s.TotalPoints == 0 {
		s.CalcTotalPoints()
	}

	s.AveragePoints = float64(s.TotalPoints) / float64(len(s.Location.Users))
}

func (s *SavedGame) CalcTotalPoints(){
	for _, user := range s.Location.Users {
		s.TotalPoints += user.Score
	}
}
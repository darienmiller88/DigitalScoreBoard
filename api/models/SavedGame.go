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
	Location      *Location          `bson:",omitempty"     json:"location"`
	Teams         *[]Team            `bson:",omitempty"     json:"teams"`
	TotalPoints   int 			     `bson:"total_points"   json:"total_points"`
	AveragePoints float64            `bson:"average_points" json:"average_points"`
	Winner        UserCard           `bson:"winner"         json:"winner"`
}

func (s *SavedGame) Validate() error{
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Winner, validation.By(s.findWinner)),
		validation.Field(&s.Location, validation.By(s.validateLocation)),
		validation.Field(&s.Teams, validation.By(s.validateTeams)),
	)
}

func (s *SavedGame) InitCreatedAtAndUpdatedAt(){
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	s.Location.InitCreatedAtAndUpdatedAt()
}

func (s *SavedGame) validateLocation(field interface{}) error{
	location, ok := field.(*Location)

	if !ok {
		return fmt.Errorf("could not parse %T into object", field)
	}

	if location == nil && s.Teams == nil {
		return fmt.Errorf("fields 'location' and 'teams' both cannot be empty")
	}

	if location == nil && s.Teams != nil {
		return fmt.Errorf("fields 'teams' requires field 'location'")
	}

	//User can pass in a non-nil Location and a nil team because this means a single player game was played.
	return s.Location.Validate()
}

func (s *SavedGame) validateTeams(field interface{}) error{
	teamLimit := 2

	if s.Teams != nil && len(*s.Teams) < teamLimit {
		return fmt.Errorf("please include at least %d teams", teamLimit)
	}

	return nil
}

func (s *SavedGame) findWinner(field interface{}) error{
	winner, ok := field.(UserCard)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}

	//Try to Find the winner only when a single player game is played, not a team game. If the winner of the game
	//was not found in the array of people at the ADAPT location, return the error.
	if (s.Teams == nil && s.Location != nil) && !slices.Contains(s.Location.Users, winner){
		return fmt.Errorf("could not find winner '%s' in list of players", winner.Name)
	}

	return nil
}

func (s *SavedGame) CalcAveragePoints(){
	//If there are either no people added to the ADAPT location yet, do not calculate the average.
	if len(s.Location.Users) == 0{
		return
	}

	//If there are no teams, do not calculate the average.
	if s.Teams != nil && len(*s.Teams) == 0 {
		return
	}

	if s.TotalPoints == 0 {
		s.CalcTotalPoints()
	}

	//If the user is playing a team game, calculate the average points using the number of teams, otherwise calculate it
	//using the total number of people at each ADAPT location.
	if s.Teams != nil {
		s.AveragePoints = float64(s.TotalPoints) / float64(len(*s.Teams))
	}else{
		s.AveragePoints = float64(s.TotalPoints) / float64(len(s.Location.Users))
	}
}

func (s *SavedGame) CalcTotalPoints(){
	//If the user is playing a team game, calculate the total points using the teams, otherwise calculate it
	//using each player at location.
	if s.Teams != nil {
		for _, team := range *s.Teams {
			s.TotalPoints += team.Score
		}
	} else {
		for _, user := range s.Location.Users {
			s.TotalPoints += user.Score
		}
	}
}
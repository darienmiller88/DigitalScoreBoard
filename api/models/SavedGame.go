package models

import (
	"fmt"
	"time"

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

		//Order matters here! Validate the location first to ensure that the Location field is not nil.
		validation.Field(&s.Location, validation.By(s.validateLocation)),
		
		//Afterwards, either field can be validated. I'm choosing to validate the winner field.
		// validation.Field(&s.Winner, validation.By(s.findWinner)),

		//Finally, validate the teams field if the user chooses to add it.
		validation.Field(&s.Teams, validation.By(s.validateTeams)),
	)
}

func (s *SavedGame) InitCreatedAtAndUpdatedAt(){
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	s.Location.InitCreatedAtAndUpdatedAt()
}

func (s *SavedGame) FindWinner() UserCard{
	if s.Teams != nil {
		teams := *s.Teams
		winningTeam := teams[0]

		for _, team := range teams{
			if team.Score > winningTeam.Score {
				winningTeam = team
			}
		}

		s.Winner.Name = winningTeam.TeamName
		s.Winner.Score = winningTeam.Score
	} else{
		players := s.Location.Users
		winningPlayer := players[0]

		for _, player := range players{
			if player.Score > winningPlayer.Score{
				winningPlayer = player
			}
		}

		s.Winner = winningPlayer
	}

	return s.Winner
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
	if s.Teams != nil {
		teamLimit := 2
		
		//If the client includes the Teams field, ensure they include exactly 2.
		if len(*s.Teams) != teamLimit {
			return fmt.Errorf("please include at only %d teams", teamLimit)
		} 

		uniqueTeams := make(map[string]int)

		//In order to see if there are duplciate teams, create a map out of the teams the client sent.
		for _, team := range *s.Teams{
			uniqueTeams[team.TeamName] = 0
		}

		//If the number of unique teams is less than the total number of teams, there are duplicates.
		if len(uniqueTeams) < len(*s.Teams) {
			return fmt.Errorf("no duplicate teams allowed")
		}

		//if they include the Teams field, and it has only 2 unique teams, validate each team to ensure each team is valid,
		//which entails a valid ADAPT location, and actual people there.
		for _, team := range *s.Teams {
			if err := team.Validate(); err != nil {
				return err
			}
		}
	} 

	return nil
}

func (s *SavedGame) CalcAveragePoints(){
	//If there are either no people added to the ADAPT location yet, do not calculate the average.
	if len(s.Location.Users) == 0{
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
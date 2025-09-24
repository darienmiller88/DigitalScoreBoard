package models

import (
	"DigitalScoreBoard/api/database"
	"context"
	"fmt"
	"math"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SavedGame struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"  json:"id"` // The MongoDB document ID
	CreatedAt     time.Time          `bson:"created_at"     json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"     json:"updated_at"`
	LocationName  string			 `bson:"location_name"  json:"location_name"`
	Players       *[]UserCard        `bson:",omitempty"     json:"players,omitempty"`
	Teams         *[]Team            `bson:",omitempty"     json:"teams,omitempty"`
	TotalPoints   int 			     `bson:"total_points"   json:"total_points"`
	AveragePoints float64            `bson:"average_points" json:"average_points"`
	Winner        UserCard           `bson:"winner"         json:"winner"`
}

func (s *SavedGame) Validate() error{
	return validation.ValidateStruct(
		s,

		//Order matters here! Validate the location name first to ensure that the Location field is not nil.
		validation.Field(&s.LocationName, validation.Required, validation.By(s.validateLocationName)),

		//Validate the teams field if the user chooses to add it.
		validation.Field(
			&s.Teams, 

			//Validate to check to make sure that the client did not send in both a non-nil players AND teams field
			validation.By(s.validatePlayersAndTeams), 

			//Afterwards, enforce a "Required" requirement for the teams field when there is no Players field
			validation.Required.When(s.Players == nil), 

			//When both checks pass, validate each team the client sent 
			validation.By(s.validateTeams),
		),
		
		//Validate the players field if the user chooses to add it.
		validation.Field(
			&s.Players, 

			//Validate to check to make sure that the client did not send in both a non-nil players AND teams field
			validation.By(s.validatePlayersAndTeams), 

			//Afterwards, enforce a requirement for the players field when there is no Teams field
			validation.Required.When(s.Teams == nil).Error("Must include at least one player"), 

			//When both checks pass, validate each player the client sent 
			validation.By(s.validatePlayers),
		),
	)
}

func (s *SavedGame) InitCreatedAtAndUpdatedAt(){
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *SavedGame) CalcAveragePoints(){
	//If the user is playing a team game, calculate the average points using the number of teams, otherwise calculate it
	//using the total number of people at each ADAPT location.
	if s.Teams != nil {
		s.AveragePoints = math.Round(float64(s.TotalPoints) / float64(len(*s.Teams)))
	}else{
		s.AveragePoints = math.Round(float64(s.TotalPoints) / float64(len(*s.Players)))
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
		for _, user := range *s.Players {
			s.TotalPoints += user.Score
		}
	}
}

func (s *SavedGame) CalculateWinner() {
	if s.Teams != nil {
		teams := *s.Teams
		winningTeam := teams[0]

		//Set the first team as the winning team, and compare each subsequent team to see whose score is the highest
		for _, team := range teams[1:]{
			if team.Score > winningTeam.Score {
				winningTeam = team
			}
		}

		s.Winner.Name = winningTeam.TeamName
		s.Winner.Score = winningTeam.Score
	} else{
		players := *s.Players
		winningPlayer := players[0]

		//Set the first player as the winning player, and compare each subsequent player to see whose score is the highest
		for _, player := range players[1:]{
			if player.Score > winningPlayer.Score{
				winningPlayer = player
			}
		}

		s.Winner = winningPlayer
	}
}

//Ensure that players and teams aren't both included
func (s *SavedGame) validatePlayersAndTeams(field interface{}) error{
	if s.Players != nil && s.Teams != nil {
		return fmt.Errorf("field 'players' and field 'teams' both cannot be included")
	}

	return nil
}

//Validate each player in the array of players to gurauntee they all exist for the location
func (s *SavedGame) validatePlayers(field interface{}) error{
	if s.Players != nil {
		locations, err := getLocations()

		if err != nil {
			return err
		}

		players := []UserCard{}

		//Extract all of the players from all of the locations
		for _, location := range locations {
			players = append(players, location.Users...)
		}

		uniquePlayerNames := make(map[string]int)

		//Add each player to a map for easier indexing when comparing the list players sent here by the client.
		for _, player := range players{
			uniquePlayerNames[player.Name] = 0
		}

		//Check to see if any player in the list of players sent by the client exists.
		for _, player := range *s.Players{
			if _, exists := uniquePlayerNames[player.Name]; !exists {
				return fmt.Errorf("player '%s' does not exist", player.Name)
			}
		}
	}	

	return  nil
}

func (s *SavedGame) validateLocationName(field interface{}) error {
	locationName, ok := field.(string)

	if !ok {
		return fmt.Errorf("could not parse %T into object", field)	
	}

	_, err := getLocationByName(locationName)

	if err != nil {
		return err
	}

	return nil
}

func (s *SavedGame) validateTeams(field interface{}) error{
	if s.Teams != nil {
		teamLimit := 2
		
		//If the client includes the Teams field, ensure they include exactly 2.
		if len(*s.Teams) != teamLimit {
			return fmt.Errorf("please include exactly %d teams", teamLimit)
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

func getAllPlayers() ([]Location, error){
	locationsCollection := database.GetLocationsCollection()
	result, err := locationsCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return []Location{}, err
	}

	locations := []Location{}

	if err := result.All(context.Background(), &locations); err != nil {
		return []Location{}, err
	}

	return locations, nil
}

package models

import (
	"DigitalScoreBoard/api/database"
	"context"
	"fmt"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
)

type Team struct{
	Score     int      `bson:"score"      json:"score"` 
	TeamName  string   `bson:"team_name"  json:"team_name"`
	Players   []string `bson:"players"    json:"players"`
}

func (t *Team) Validate() error{
	return validation.ValidateStruct(
		t,
		validation.Field(&t.TeamName, validation.By(t.checkTeamNameInLocations)),
		// validation.Field(&t.Players, validation.By(t.checkTeamPlayers)),
	)
}

// func (t *Team) checkTeamPlayers(field interface{}) error {
// 	players, ok := field.([]string)

// 	if !ok{
// 		return fmt.Errorf("could not parse %T into object", field)
// 	}

// 	locations, err := getLocations()

// 	if err != nil{
// 		return err
// 	}
	
// 	chosenLocationWithTeamName := Location{}

// 	for _, location := range locations {
// 		if t.TeamName == location.LocationName {
// 			chosenLocationWithTeamName = location
// 			break
// 		}
// 	}

// 	slices.Index()

// }

func (t *Team) checkTeamNameInLocations(field interface{}) error{
	locationName, ok := field.(string)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}

	locations, err := getLocations()

	if err != nil{
		return err
	}

	for _, location := range locations {
		if locationName == location.LocationName {
			return nil
		}
	}

	locationNames := make([]string, len(locations)) 

	for i, location := range locations {
		locationNames[i] = location.LocationName
	}	

	return fmt.Errorf("Team name '%s' is not a valid team name. Please choose from " +
		"the following: [%s] ", locationName, strings.Join(locationNames, ", "))
}

func getLocations() ([]Location, error){
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
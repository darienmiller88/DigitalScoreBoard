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

// Check to see if the team name passed in from the client actually exists as a Valid ADAPT Community Location.
// In this case, the "team name" is just the name of the ADAPT location, not a customizable team name.
func (t *Team) checkTeamNameInLocations(field interface{}) error{
	locationName, ok := field.(string)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}

	locations, err := getLocations()

	if err != nil{
		return err
	}

	//Retrieve all of the locations from the database.
	for _, location := range locations {
		if locationName == location.LocationName {
			return nil
		}
	}

	locationNames := make([]string, len(locations)) 

	//Sadly, there is no .Map() function to extract the location name from the Location object into a array of strings.
	for i, location := range locations {
		locationNames[i] = location.LocationName
	}	

	//Finally, return a error message signaling to the user that the team name MUST be one of the valid ADAPT locations.
	return fmt.Errorf("Team name '%s' is not a valid team name. Please choose from " +
		"the following: [%s] ", locationName, strings.Join(locationNames, ", "))
}

//Function to retrieve all ADAPT locations from database. Sadly, I cannot use the service for this as that would result
// in a circular depency :(
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
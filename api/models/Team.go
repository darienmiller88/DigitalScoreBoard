package models

import (
	"DigitalScoreBoard/api/database"
	"context"
	"fmt"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		validation.Field(&t.Players, validation.Required, validation.Length(2, 0), validation.By(t.checkDuplicateTeamPlayers), validation.By(t.checkForValidTeamPlayers)),
	)
}

//Check for duplicates in the list of users sent by the client.
func (t *Team) checkDuplicateTeamPlayers (field interface{}) error{
	players, ok := field.([]string)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}
	
	//Create a map of string to int (value is irrelevant, I only care about the key)
	uniqueNames := make(map[string]int)

	//In order to see how if there are duplicate names, add all of the names the client passed into the map I just made.
	//Let's say players len = 5. If all 5 names are unique, that map will also have 5 elements. If there is one
	//duplicate (m, n, k, l, l), the map will have a length of 4 because the second "l" in that example is not given 
	//its own bucket in the map, the value of the first "l" is simply overridden.
	for _, player := range players{
		uniqueNames[player] = 0
	}
	
	//When the map is made, check to see if the len of the players array is greater than the map. If so, 
	//then there are duplicates!
	if len(players) > len(uniqueNames) {
		return fmt.Errorf("no duplicate names allowed! [%s]", strings.Join(players, ", "))
	}

	//If not, return nil signifying validation success.
	return nil
}

//Check to see if any of the players the client sent are invalid people at the specific ADAPT location. 
func (t *Team) checkForValidTeamPlayers(field interface{}) error {
	players, ok := field.([]string)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}

	//Since the team name is validated first to ensure it matches a valid ADAPT location, we can just send it the 
	//below method to get the Location object by the team name.
	location, err := getLocationByName(t.TeamName)

	if err != nil{
		return err
	}

	//Save the amount of people in the players array sent by the client, as well as the number of people at the 
	//actual ADAPT location.
	numPlayers          := len(players)
	numPeopleAtLocation := len(location.Users) 
	
	//If the client sent more players than actual people at this ADAPT location, send an error to prevent time wasting.
	if numPlayers > numPeopleAtLocation {
		return fmt.Errorf("length of players field cannot exceed total amount of people at ADAPT location: %d > %d", numPlayers, numPeopleAtLocation)
	}

	uniquePlayersMap := make(map[string]int)
	
	//First create a map out of the players at a given adapt location to allow constant time lookup for each person.
	for _, player := range location.Users{
		uniquePlayersMap[player.Name] = 0
	}

	//Afterwards, check each player the client sent to see if it exists in the map. If not, return the following error.
	for _, newlyAddedplayer := range players{
		if _, exists := uniquePlayersMap[newlyAddedplayer]; !exists {
			return fmt.Errorf("'%s' is not a valid person at the ADAPT location '%s'", newlyAddedplayer, t.TeamName)
		}
	}
	
	//If every name the client sent matched the names of the players in the ADAPT location, return nil to allow
	//validation success.
	return nil
}

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

		//If the location name sent by the client matches a Valid ADAPT location, return nil as the field is valid.
		if locationName == location.LocationName {
			return nil
		}
	}

	//If the location did NOT match any of the ADAPT location names, prepare this data to send back to them.
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
// in a circular dependency :(
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

func getLocationByName(locationName string) (Location, error){
	location := Location{}
	locationsCollection := database.GetLocationsCollection()
	err := locationsCollection.FindOne(context.Background(),  bson.D{{Key: "location_name", Value: locationName}}).Decode(&location)

	if err != nil {
		if err == mongo.ErrNoDocuments{
			return Location{}, fmt.Errorf("Location '%s' not a valid location", locationName)
		}else{
			return Location{}, err
		}
	}

	return location, nil
}
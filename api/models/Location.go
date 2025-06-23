package models

import (
	"context"
	"fmt"
	"time"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"DigitalScoreBoard/api/database"
)

// This struct will represent the ADAPT community network location where a jeopardy game was played.
type Location struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"` 
	CreatedAt    time.Time          `bson:"created_at"    json:"created_at"` 
	UpdatedAt    time.Time          `bson:"updated_at"    json:"updated_at"`

	//The name of the ADAPT location (Pelham Bay, Lawrence, Elmwood, etc)
	LocationName string             `bson:"location_name" json:"location_name"`

	//Here are all of the users that played in the game.
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
		validation.Field(&l.LocationName, validation.Required, validation.By(l.findLocation)),
	)
}

//Check to see if the location name is a valid location name based on a pre-defined set I have in mongoDB.
func (l *Location) findLocation(field interface{}) error{
	locationName, ok := field.(string)

	if !ok{
		return fmt.Errorf("could not parse %T into object", field)
	}

	locationsCollection := database.GetLocationsCollection()
	result, err := locationsCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return err
	}

	locations := []Location{}

	if err := result.All(context.Background(), &locations); err != nil {
		return err
	}

	//Check to see if there's a match between the location name the user sent, and the ones in the database.
	for _, location := range locations{
		if locationName == location.LocationName {
			return nil
		}
	}

	locationNames := make([]string, len(locations)) 

	//If there's no match, extract the location names from the Location objects (sadly there's no .Map() 
	// function to do this...) to send back as part of the error message.
	for i, location := range locations {
		locationNames[i] = location.LocationName
	}	

	return fmt.Errorf("Location name '%s' is not a valid location name. Please choose from " +
		"the following: [%s]", l.LocationName, strings.Join(locationNames, ", "))
}
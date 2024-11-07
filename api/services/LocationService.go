package services

import (
	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Retrieve all locations from database.
func GetAllLocations(req *http.Request) ([]models.Location, error) {
	locationsCollection := database.GetLocationsCollection()
	result, err := locationsCollection.Find(req.Context(), bson.D{})

	if err != nil {
		return []models.Location{}, err
	}

	locations := []models.Location{}

	if err := result.All(req.Context(), &locations); err != nil {
		return []models.Location{}, err
	}

	return locations, nil
}

//Retrieve one location from MongoDB.
func GetLocation(req *http.Request, locationName string) models.LocationResult{
	locationsCollection := database.GetLocationsCollection()
	location := &models.Location{}
	result := models.LocationResult{}

	err := locationsCollection.FindOne(req.Context(), bson.D{{Key: "location_name", Value: locationName}}).Decode(&location)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			result.StatusCode = http.StatusNotFound
		} else {
			result.StatusCode = http.StatusInternalServerError
		}

		result.Err = err
		return result
	}

	result.Location = *location
	result.StatusCode = http.StatusOK

	return result
}

//Service function to allow adding a user to an Adapt location, and removing.
func UpdateUsersForLocation(req *http.Request, mongoUpdateOperator string, locationName string, username string) (*mongo.UpdateResult, error){
	// if mongoUpdateOperator != "$pull"  ||{
		
	// }
	
	filter := bson.M{"location_name": locationName}
	update := bson.M{mongoUpdateOperator: bson.M{"users": bson.M{"name": username}}}

 	result, err := database.GetLocationsCollection().UpdateOne(req.Context(), filter, update)

	if err != nil{
		return nil, err
	}

	return result, nil
}
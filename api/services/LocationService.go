package services

import (
	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const(
	PUSH string = "$push"
	PULL string = "$pull"
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
func GetLocation(req *http.Request, locationName string) models.Result[models.Location]{
	locationsCollection := database.GetLocationsCollection()
	location            := &models.Location{}
	result              := models.Result[models.Location]{}
	err                 := locationsCollection.FindOne(
		req.Context(), 
		bson.D{{Key: "location_name", Value: locationName}},
	).Decode(&location)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			result.StatusCode = http.StatusNotFound
		} else {
			result.StatusCode = http.StatusInternalServerError
		}

		result.Err = err
		return result
	}

	result.ResultData = *location
	result.StatusCode = http.StatusOK

	return result
}

//Service function to allow adding or removing a user to or from an Adapt location.
func UpdateUsersForLocation(req *http.Request, mongoUpdateOperator string, locationName string, username string) models.Result[*mongo.UpdateResult] {
	updateUserResult := models.Result[*mongo.UpdateResult]{
		StatusCode: http.StatusInternalServerError,
	}
	
	if !(mongoUpdateOperator == PULL || mongoUpdateOperator == PUSH){
		updateUserResult.Err = fmt.Errorf("mongo operator %s not valid. Must be either %s or %s", mongoUpdateOperator, PULL, PUSH)

		return updateUserResult
	}
	

	filter := bson.M{"location_name": locationName}
	update := bson.M{mongoUpdateOperator: bson.M{"users": bson.M{"name": username}}}
 	updateOneResult, err := database.GetLocationsCollection().UpdateOne(req.Context(), filter, update)

	if err != nil{
		updateUserResult.Err = err
		
		return updateUserResult
	}

	if updateOneResult.ModifiedCount == 0 {
		updateUserResult.Err = fmt.Errorf("no location \"%s\" found", locationName)
		updateUserResult.StatusCode = http.StatusNotFound

		return updateUserResult
	}

	updateUserResult.ResultData = updateOneResult
	updateUserResult.StatusCode = http.StatusOK

	return updateUserResult
}
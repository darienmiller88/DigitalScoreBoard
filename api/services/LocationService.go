package services

import (
	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const(
	PUSH string = "$push"
	PULL string = "$pull"
)

//Add a new adapt location, which for now, I will not expose.
func AddNewAdaptLocation(req *http.Request, location models.Location) models.Result[models.Location]{
	insertOneResult, err := database.GetLocationsCollection().InsertOne(req.Context(), location)
	locationResult := models.Result[models.Location]{}

	if err != nil{
		locationResult.Err = err
		locationResult.StatusCode = http.StatusInternalServerError

		return locationResult
	}
	
	location.ID = insertOneResult.InsertedID.(primitive.ObjectID)
	locationResult.ResultData = location
	locationResult.StatusCode = http.StatusOK

	return locationResult
}

//Retrieve all locations from database.
func GetAllLocations(req *http.Request) models.Result[[]models.Location] {
	locationsCollection := database.GetLocationsCollection()
	findResult, err := locationsCollection.Find(req.Context(), bson.D{})

	if err != nil {
		return models.Result[[]models.Location]{ StatusCode: http.StatusInternalServerError, Err: err }
	}

	locations := []models.Location{}

	if err := findResult.All(req.Context(), &locations); err != nil {
		return models.Result[[]models.Location]{ StatusCode: http.StatusInternalServerError, Err: err }
	}

	return models.Result[[]models.Location]{ StatusCode: http.StatusOK, ResultData: locations }
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

func UpdatePlayerName(req *http.Request, locationName string, playerName string) models.Result[*mongo.UpdateResult]{
	updateUserResult := models.Result[*mongo.UpdateResult]{
		StatusCode: http.StatusInternalServerError,
	}
	
	// filter := bson.M{"location_name": locationName}
	// update := bson.M{"$set": bson.M{"users": bson.M{"name": playerName}}}

	// updateOneResult, err := database.GetLocationsCollection().UpdateOne(req.Context(), filter, update)


	return updateUserResult
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
		updateUserResult.Err = fmt.Errorf("no location \"%s\" found OR player \"%s\" not found", locationName, username)
		updateUserResult.StatusCode = http.StatusNotFound

		return updateUserResult
	}

	updateUserResult.ResultData = updateOneResult
	updateUserResult.StatusCode = http.StatusOK

	return updateUserResult
}
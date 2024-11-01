package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"DigitalScoreBoard/api/utilities"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type locationResult struct{
	statusCode int
	location   models.Location
	err        error
}

// please so change
//Retrieve one location from the list of Adapt locations.
func GetLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	result := getLocation(req, locationName)

	if result.err != nil{
		http.Error(res, result.err.Error(), result.statusCode)
		return
	}
	
	//  
	// 
	utilities.SendJSON(http.StatusOK, res, result.location)
}

func GetAllUsersByLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	result := getLocation(req, locationName)

	if result.err != nil {
		http.Error(res, result.err.Error(), result.statusCode)
		return
	}

	utilities.SendJSON(http.StatusOK, res, result.location.Users)
}

//Get all of the Adapt locations.
func GetAllLocations(res http.ResponseWriter, req *http.Request){
	locationsCollection := database.GetDB().Database(database.DatabaseName).Collection(database.LocationsCollection)
	result, err := locationsCollection.Find(req.Context(), bson.D{})

	if err != nil {
		utilities.SendJSON(http.StatusInternalServerError, res, err.Error())
		return
	}

	locations := []models.Location{}

	if err := result.All(req.Context(), &locations); err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	utilities.SendJSON(http.StatusOK, res, locations)
}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	
}

func GetAllSavedGames(res http.ResponseWriter, req *http.Request){
	savedGamesCollection := database.GetDB().Database(database.DatabaseName).Collection(database.SavedGamesCollection)
	result, err := savedGamesCollection.Find(req.Context(), bson.D{})

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	savedGames := []models.SavedGame{}

	if err := result.All(req.Context(), &savedGames); err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}


	utilities.SendJSON(http.StatusOK, res, savedGames)
}

func SaveGame(res http.ResponseWriter, req *http.Request){
	data := make(map[string]interface{})
	
	json.NewDecoder(req.Body).Decode(&data)
	fmt.Println(data)

	utilities.SendJSON(http.StatusOK, res, data)
}

func AddLocation(res http.ResponseWriter, req *http.Request){
	location := models.Location{}

	if err := json.NewDecoder(req.Body).Decode(&location); err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := location.Validate(); err != nil{
		utilities.SendJSON(http.StatusBadRequest, res, err)
		return
	}

	location.InitCreatedAtAndUpdatedAt()
	result, err := database.GetDB().Database(database.DatabaseName).Collection(database.LocationsCollection).InsertOne(req.Context(), location)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	location.ID = result.InsertedID.(primitive.ObjectID)
	utilities.SendJSON(http.StatusOK, res, location)
}

func AddUserToLocation(res http.ResponseWriter, req *http.Request){
	location := chi.URLParam(req, "location-name")
	username := struct{
		Username string `json:"username"`
	}{}

	if err := json.NewDecoder(req.Body).Decode(&username); err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := updateUsersForLocation(req, "$push", location, username.Username)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.ModifiedCount == 0 {
		http.Error(res, fmt.Sprintf("No location \"%s\" found", location), http.StatusNotFound)
		return
	}

	utilities.SendJSON(http.StatusOK, res, map[string]interface{}{
		"message": fmt.Sprintf("%s was added to location %s", username.Username, location),
		"result": result,
	})
}

func RemoveUserFromLocation(res http.ResponseWriter, req *http.Request){
	location := chi.URLParam(req, "location-name")
	username := chi.URLParam(req, "username")
	result, err := updateUsersForLocation(req, "$pull", location, username)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	
	if result.ModifiedCount == 0 {
		http.Error(res, fmt.Sprintf("No user \"%s\" or location %s found", username, location), http.StatusNotFound)
		return
	}

	utilities.SendJSON(http.StatusOK, res, map[string]interface{}{
		"message": fmt.Sprintf("%s was removed from location %s", username, location),
		"result": result,
	})
}

//Utility function to allow adding a user to an Adapt location, and removing 
func updateUsersForLocation(req *http.Request, mongoUpdateOperator string, locationName string, username string) (*mongo.UpdateResult, error){
	filter := bson.M{"location_name": locationName}
	update := bson.M{mongoUpdateOperator: bson.M{"users": bson.M{"name": username}}}

 	result, err := database.GetDB().
					Database(database.DatabaseName).
					Collection(database.LocationsCollection).UpdateOne(req.Context(), filter, update)

	if err != nil{
		return nil, err
	}

	return result, nil
}
 
//Utility function to retrieve location from MongoDB.
func getLocation(req *http.Request, locationName string) locationResult{
	locationsCollection := database.GetDB().Database(database.DatabaseName).Collection(database.LocationsCollection)
	location := &models.Location{}
	result := locationResult{}

	err := locationsCollection.FindOne(req.Context(), bson.D{{Key: "location_name", Value: locationName}}).Decode(&location)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			result.statusCode = http.StatusNotFound
		} else {
			result.statusCode = http.StatusInternalServerError
		}

		result.err = err
		return result
	}

	result.location = *location
	result.statusCode = http.StatusOK

	return result
}

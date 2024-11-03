package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"DigitalScoreBoard/api/utilities"
	"DigitalScoreBoard/api/services"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Retrieve one location from the list of Adapt locations.
func GetLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	result := services.GetLocation(req, locationName)

	if result.Err != nil{
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}
	
	utilities.SendJSON(http.StatusOK, res, result.Location)
}

func GetAllUsersByLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	result := services.GetLocation(req, locationName)

	if result.Err != nil {
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}

	utilities.SendJSON(http.StatusOK, res, result.Location.Users)
}

//Get all of the Adapt locations.
func GetAllLocations(res http.ResponseWriter, req *http.Request){
	locations, err := services.GetAllLocations(req)

	if err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	utilities.SendJSON(http.StatusOK, res, locations)
}

//Get all users for ALL locations.
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	locations, err := services.GetAllLocations(req)

	if err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	users := []models.UserCard{}

	for _, location := range locations {
		users = append(users, location.Users...)
	}

	utilities.SendJSON(http.StatusOK, res, users)
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

	result, err := services.UpdateUsersForLocation(req, "$push", location, username.Username)

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
	result, err := services.UpdateUsersForLocation(req, "$pull", location, username)

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
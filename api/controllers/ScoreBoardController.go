package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"DigitalScoreBoard/api/utilities"
	"DigitalScoreBoard/api/services"

	"github.com/go-chi/chi/v5"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

const(
	PULL string = "$pull"
	PUSH string = "$push"
)

//Retrieve one location from the list of Adapt locations.
func GetLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	result := services.GetLocation(req, locationName)

	if result.Err != nil{
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}
	
	utilities.SendJSON(http.StatusOK, res, result.ResultData)
}

func GetAllUsersByLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	result := services.GetLocation(req, locationName)

	if result.Err != nil {
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}

	utilities.SendJSON(http.StatusOK, res, result.ResultData.Users)
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

//Get ALL users for ALL locations.
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

//Retrieve all saved game from ONE location.
func GetAllSavedGamesFromLocation(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	 
	//Retrieve games from service.
	savedGamesResult := services.GetAllSavedGamesFromLocation(req, locationName)

	if savedGamesResult.Err != nil {
		http.Error(res, savedGamesResult.Err.Error(), savedGamesResult.StatusCode)
		return
	}

	utilities.SendJSON(savedGamesResult.StatusCode, res, savedGamesResult.ResultData)
}

//Retrieve all saved game for ALL locations.
func GetAllSavedGames(res http.ResponseWriter, req *http.Request){
	savedGamesResult := services.GetAllSavedGames(req)

	if savedGamesResult.Err != nil {
		http.Error(res, savedGamesResult.Err.Error(), savedGamesResult.StatusCode)
		return
	}

	utilities.SendJSON(http.StatusOK, res, savedGamesResult.ResultData)
}

//Save a game to the database.
func SaveGame(res http.ResponseWriter, req *http.Request){
	savedGame := models.SavedGame{}

	if err := json.NewDecoder(req.Body).Decode(&savedGame); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("savedGame:", savedGame)

	if err := savedGame.Validate(); err != nil{
		utilities.SendJSON(http.StatusBadRequest, res, err)
		return
	}

	// if err := savedGame.Validate(); err != nil{
	// 	savedGamesResult.Err = err
	// 	savedGamesResult.StatusCode = http.StatusBadRequest

	// 	return savedGamesResult
	// }
	// result := services.AddSavedGame(req, savedGame)


	// if result.Err != nil {
	// 	http.Error(res, result.Err.Error(), result.StatusCode)
	// 	return
	// }

	utilities.SendJSON(http.StatusOK, res, savedGame)
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
	// result, err := database.GetLocationsCollection().InsertOne(req.Context(), location)

	// if err != nil {
	// 	http.Error(res, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// location.ID = result.InsertedID.(primitive.ObjectID)
	utilities.SendJSON(http.StatusOK, res, location)
}

func UpdateScore(res http.ResponseWriter, req *http.Request){

}

func AddUserToLocation(res http.ResponseWriter, req *http.Request){
	modifyUserInLocation(res, req, services.PUSH)
}

func RemoveUserFromLocation(res http.ResponseWriter, req *http.Request){
	modifyUserInLocation(res, req, services.PULL)
}

func modifyUserInLocation(res http.ResponseWriter, req *http.Request, operation string){
	location := chi.URLParam(req, "location-name")
	username := struct{
		Username string `json:"username"`
	}{}

	// Decode the username from request body
	if err := json.NewDecoder(req.Body).Decode(&username); err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.UpdateUsersForLocation(req, operation, location, username.Username)

	if result.Err != nil {
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}

	message := "added to"
	if operation == services.PULL  {
		message = "removed from"
	}

	utilities.SendJSON(http.StatusOK, res, map[string]interface{}{
		"message": fmt.Sprintf("%s was %s location: %s", username, message, location),
		"result": result,
	})
}
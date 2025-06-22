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
	var locationsResult models.Result[[]models.Location] = services.GetAllLocations(req)

	if locationsResult.Err != nil{
		http.Error(res, locationsResult.Err.Error(), http.StatusInternalServerError)
		return
	}

	utilities.SendJSON(http.StatusOK, res, locationsResult.ResultData)
}

//Get ALL users for ALL locations.
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	var result models.Result[[]models.Location] = services.GetAllLocations(req)

	if result.Err != nil{
		http.Error(res, result.Err.Error(), http.StatusInternalServerError)
		return
	}

	users := []models.UserCard{}

	for _, location := range result.ResultData {
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

func UpdatePlayerName(res http.ResponseWriter, req *http.Request){
	location := chi.URLParam(req, "location-name")
	playerName := struct{
		PlayerName string `json:"player_name"`
	}{}

	// Decode the player name from request body
	if err := json.NewDecoder(req.Body).Decode(&playerName); err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	locationResult := services.GetLocation(req, location)

	if locationResult.Err != nil {
		http.Error(res, locationResult.Err.Error(), locationResult.StatusCode)
		return
	}

	//Using both the location and new playername, call the following service and provide both to it.
	updateResult := services.UpdatePlayerName(req, location, playerName.PlayerName)

	if updateResult.Err != nil {
		http.Error(res, updateResult.Err.Error(), updateResult.StatusCode)
		return
	}

	utilities.SendJSON(updateResult.StatusCode, res, utilities.M{
		"message": fmt.Sprintf("player name updated %s", playerName.PlayerName),
		"update_result": updateResult,
	})
}

func AddUserToLocation(res http.ResponseWriter, req *http.Request){
	modifyUserInLocation(res, req, services.PUSH)
}

func RemoveUserFromLocation(res http.ResponseWriter, req *http.Request){
	modifyUserInLocation(res, req, services.PULL)
}

func modifyUserInLocation(res http.ResponseWriter, req *http.Request, operation string){
	location := chi.URLParam(req, "location-name")
	playerName := struct{
		PlayerName string `json:"player_name"`
	}{}

	// Decode the username from request body
	if err := json.NewDecoder(req.Body).Decode(&playerName); err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.UpdateUsersForLocation(req, operation, location, playerName.PlayerName)

	if result.Err != nil {
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}

	message := "added to"
	if operation == services.PULL  {
		message = "removed from"
	}

	utilities.SendJSON(http.StatusOK, res, utilities.M{
		"message": fmt.Sprintf("%s was %s location: %s", playerName.PlayerName, message, location),
		"result": result,
	})
}
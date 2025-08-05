package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	// "DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"DigitalScoreBoard/api/services"
	"DigitalScoreBoard/api/utilities"

	"github.com/go-chi/chi/v5"
	"github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func DeleteSavedGame(res http.ResponseWriter, req *http.Request){
	id, err := primitive.ObjectIDFromHex(chi.URLParam(req, "id"))

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var deleteResult models.Result[*mongo.DeleteResult] = services.DeleteSavedGame(req, id)

	if deleteResult.Err != nil{
		http.Error(res, deleteResult.Err.Error(), deleteResult.StatusCode)
		return
	}

	utilities.SendJSON(http.StatusOK, res, "Saved game successfully deleted!")
}

//Save a game to the database.
func SaveGame(res http.ResponseWriter, req *http.Request){
	savedGame := models.SavedGame{}

	//Decode the body into a SavedGame object
	if err := json.NewDecoder(req.Body).Decode(&savedGame); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Try to validate the saved game
	if err := savedGame.Validate(); err != nil{
		utilities.SendJSON(http.StatusBadRequest, res, err)
		return
	}

	var result models.Result[models.SavedGame] = services.AddSavedGame(req, savedGame)

	if result.Err != nil {
		http.Error(res, result.Err.Error(), result.StatusCode)
		return
	}

	utilities.SendJSON(result.StatusCode, res, result.ResultData)
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

//Function to allow for the editting of a players name
func UpdatePlayerName(res http.ResponseWriter, req *http.Request){
	locationName := chi.URLParam(req, "location-name")
	playerNames := struct{
		NewPlayerName string `json:"new_player_name"`
		OldPlayerName string `json:"old_player_name"`		
	}{}

	// Decode the player names from request body
	if err := json.NewDecoder(req.Body).Decode(&playerNames); err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	//Validate the above struct to ensure both fields are included, and that the new name is between 3 and 30 chars.
	err := validation.ValidateStruct(&playerNames,
		validation.Field(&playerNames.NewPlayerName, validation.Required, validation.Length(3, 31)),
		validation.Field(&playerNames.OldPlayerName, validation.Required),
	)

	//If the first round of validation does not pass, return the following error.
	if err != nil{
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	//Afterwards, get the location to ensure the client sent one that actually exists. 
	var locationResult models.Result[models.Location] = services.GetLocation(req, locationName)

	// If not, return an appropriate error.
	if locationResult.Err != nil {
		http.Error(res, locationResult.Err.Error(), locationResult.StatusCode)
		return
	}

	//Validate the old name to check to see if it actually exists the ADAPT location
	if !slices.Contains(locationResult.ResultData.Users, models.UserCard{ Name: playerNames.OldPlayerName }) {
		http.Error(res, fmt.Sprintf("Player '%s' at location '%s' does not exist", playerNames.OldPlayerName, locationName), http.StatusNotFound)
		return
	}

	//Check to see if the new name the client is trying to send is already taken.
	if slices.Contains(locationResult.ResultData.Users, models.UserCard{ Name: playerNames.NewPlayerName }) {
		http.Error(res, fmt.Sprintf("Player name '%s' is taken! Pick another one.", playerNames.OldPlayerName), http.StatusConflict)
		return
	}

	//After validating both the old player name, and the new player name, using the location call the following 
	//service and provide both names to it.
	updateResult := services.UpdatePlayerName(req, locationName, playerNames.OldPlayerName, playerNames.NewPlayerName)

	//Return the update result if an error occurred.
	if updateResult.Err != nil {
		http.Error(res, updateResult.Err.Error(), updateResult.StatusCode)
		return
	}

	//FINALLY, after all of that validating, and after the name has been updated in the database, return a success message
	utilities.SendJSON(updateResult.StatusCode, res, utilities.M{
		"message": fmt.Sprintf("player name updated from '%s' to '%s'", playerNames.OldPlayerName, playerNames.NewPlayerName),
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

	result := services.UpdatePlayersForLocation(req, operation, location, playerName.PlayerName)

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
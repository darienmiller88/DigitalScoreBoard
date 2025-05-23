package services

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
)

//This is a comment from my temp laptop to test commit success.
func GetAllSavedGames(req *http.Request) models.Result[[]models.SavedGame] {
	findResult, err := database.GetSavedGamesCollections().Find(req.Context(), bson.D{})
	
	//Initialize the result after trying to find all current saved games.
	savedGamesResult := models.Result[[]models.SavedGame]{
		StatusCode: http.StatusInternalServerError,
		Err: err,
	}

	if err != nil {
		return savedGamesResult
	}

	savedGames := []models.SavedGame{}

	if err := findResult.All(req.Context(), &savedGames); err != nil{
		savedGamesResult.Err = err
		
		return savedGamesResult
	}

	savedGamesResult.ResultData = savedGames
	savedGamesResult.StatusCode = http.StatusOK

	return savedGamesResult
}

func GetAllSavedGamesFromLocation(req *http.Request, locationName string) models.Result[[]models.SavedGame] {
	savedGames           := []models.SavedGame{}
	savedGamesCollection := database.GetSavedGamesCollections()
	err                  := savedGamesCollection.FindOne(req.Context(), bson.D{
		{Key: "location_name", Value: locationName},
	}).Decode(&savedGames)
	
	savedGamesResult := models.Result[[]models.SavedGame]{}
		
	if err != nil {
		if err == mongo.ErrNoDocuments {
			savedGamesResult.StatusCode = http.StatusNotFound
		} else {
			savedGamesResult.StatusCode = http.StatusInternalServerError
		}

		savedGamesResult.Err = err
		return savedGamesResult
	}

	savedGamesResult.StatusCode = http.StatusOK
	savedGamesResult.ResultData = savedGames

	return savedGamesResult	
}

func AddSavedGame(req *http.Request, savedGame models.SavedGame) models.Result[models.SavedGame]{
	savedGamesResult := models.Result[models.SavedGame]{}

	if err := savedGame.Validate(); err != nil{
		savedGamesResult.Err = err
		savedGamesResult.StatusCode = http.StatusBadRequest

		return savedGamesResult
	}

	savedGamesCollection := database.GetSavedGamesCollections()
	result, err          := savedGamesCollection.InsertOne(req.Context(), savedGame)

	if err != nil{ 
		savedGamesResult.Err = err
		savedGamesResult.StatusCode = http.StatusInternalServerError

		return savedGamesResult
	}

	//Get the id of the location here the game took place.
	locationId := GetLocation(req, savedGame.Location.LocationName).ResultData.ID

	//Calculate both the total amount of points earned, as well as the average.
	savedGame.CalcTotalPoints()
	savedGame.CalcAveragePoints()

	//Finally, attach the id of the location to saved game's location, and the id of the newly created 
	//saved game.
	savedGame.Location.ID = locationId
	savedGame.ID = result.InsertedID.(primitive.ObjectID)
	savedGamesResult.ResultData = savedGame
	savedGamesResult.StatusCode = http.StatusOK

	return savedGamesResult
}
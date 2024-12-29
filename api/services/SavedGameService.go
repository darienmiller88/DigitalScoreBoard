package services

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
)

func GetAllSavedGames(req *http.Request) models.Result[[]models.SavedGame] {
	savedGamesCollection := database.GetSavedGamesCollections()
	savedGamesResult     := models.Result[[]models.SavedGame]{}
	findResult, err      := savedGamesCollection.Find(req.Context(), bson.D{})
	
	savedGamesResult.StatusCode = http.StatusInternalServerError
	savedGamesResult.Err = err

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

	savedGame.ID = result.InsertedID.(primitive.ObjectID)
	savedGamesResult.ResultData = savedGame
	savedGamesResult.StatusCode = http.StatusOK

	return savedGamesResult
}
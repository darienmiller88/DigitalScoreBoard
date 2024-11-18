package services

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
)

func GetAllSavedGames(req *http.Request) models.Result[[]models.SavedGame] {
	savedGamesCollection := database.GetSavedGamesCollections()
	savedGamesResult := models.Result[[]models.SavedGame]{}
	findResult, err := savedGamesCollection.Find(req.Context(), bson.D{})

	if err != nil {
		savedGamesResult.Err = err
		savedGamesResult.StatusCode = http.StatusInternalServerError

		return savedGamesResult
	}

	savedGames := []models.SavedGame{}

	if err := findResult.All(req.Context(), &savedGames); err != nil{
		savedGamesResult.Err = err
		savedGamesResult.StatusCode = http.StatusInternalServerError
		
		return savedGamesResult
	}

	savedGamesResult.ResultData = savedGames
	savedGamesResult.StatusCode = http.StatusOK

	return savedGamesResult
}

func GetAllSavedGamesFromLocation(){

}

func AddSavedGame(){

}
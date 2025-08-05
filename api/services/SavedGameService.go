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

	//Unmarshall the mongo cursor into the array of saved games.
	if err := findResult.All(req.Context(), &savedGames); err != nil{
		savedGamesResult.Err = err
		savedGamesResult.StatusCode = http.StatusInternalServerError
		
		return savedGamesResult
	}

	savedGamesResult.ResultData = savedGames
	savedGamesResult.StatusCode = http.StatusOK

	return savedGamesResult
}

func GetAllSavedGamesFromLocation(req *http.Request, locationName string) models.Result[[]models.SavedGame] {
	savedGames           := []models.SavedGame{}
	savedGamesCollection := database.GetSavedGamesCollections()
	findResult, err      := savedGamesCollection.Find(req.Context(), bson.D{
		{ Key: "location_name", Value: locationName },
	})
	
	savedGamesResult := models.Result[[]models.SavedGame]{
		StatusCode: http.StatusInternalServerError,
	}
		
	if err != nil {
		if err == mongo.ErrNoDocuments {
			savedGamesResult.StatusCode = http.StatusNotFound
		} 

		savedGamesResult.Err = err
		return savedGamesResult
	}

	if err := findResult.All(req.Context(), &savedGames); err != nil{
		savedGamesResult.Err = err

		return savedGamesResult
	}

	savedGamesResult.StatusCode = http.StatusOK
	savedGamesResult.ResultData = savedGames

	return savedGamesResult	
}

func DeleteSavedGame(req *http.Request, savedGameId string) models.Result[*mongo.DeleteResult]{
	deleteSavedGameResult := models.Result[*mongo.DeleteResult]{}
	result, err := database.GetSavedGamesCollections().DeleteOne(req.Context(), bson.M{
		"_id": savedGameId,
	})

	if err != nil {
		deleteSavedGameResult.Err = err
		deleteSavedGameResult.StatusCode = http.StatusBadRequest

		return deleteSavedGameResult
	}

	deleteSavedGameResult.ResultData = result
	deleteSavedGameResult.StatusCode = http.StatusOK

	return deleteSavedGameResult
}

func AddSavedGame(req *http.Request, savedGame models.SavedGame) models.Result[models.SavedGame]{
	//Initialize the saved game's created and updated fields.
	savedGame.InitCreatedAtAndUpdatedAt()

	//Calculate both the total amount of points earned, as well as the average.
	savedGame.CalcTotalPoints()
	savedGame.CalcAveragePoints()
	
	//Calculate the winner of the game that was played.
	savedGame.CalculateWinner()
	
	//create the result object, get the saved games collection, and insert the saved game into the database.
	savedGamesResult     := models.Result[models.SavedGame]{}
	savedGamesCollection := database.GetSavedGamesCollections()
	insertResult, err    := savedGamesCollection.InsertOne(req.Context(), &savedGame)

	if err != nil{ 
		savedGamesResult.Err = err
		savedGamesResult.StatusCode = http.StatusInternalServerError
		
		return savedGamesResult
	}
	
	//Finally, attach the id of the newly created saved game.
	savedGame.ID = insertResult.InsertedID.(primitive.ObjectID)
	savedGamesResult.ResultData = savedGame
	savedGamesResult.StatusCode = http.StatusOK

	return savedGamesResult
}
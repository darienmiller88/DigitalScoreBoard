package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/models"
	"DigitalScoreBoard/api/utilities"

	"go.mongodb.org/mongo-driver/bson"
)

// type ScoreBoardController struct {
// }

// func(s *ScoreBoardController) Init(){

// }

func GetLocation(res http.ResponseWriter, req *http.Request){
	
}

func GetAllLocations(res http.ResponseWriter, req *http.Request){
	
}

func GetAllUsersByLocation(res http.ResponseWriter, req *http.Request){

}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	locationsCollection := database.GetDB().Database(database.DatabaseName).Collection(database.LocationsCollection)

	result, err := locationsCollection.Find(req.Context(), bson.D{})

	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	locations := []models.Location{}

	result.All(req.Context(), &locations)

	fmt.Println("locations:", locations)
	utilities.SendJSON(http.StatusOK, res, map[string]interface{}{"locations": locations})
}

func GetAllSavedGames(res http.ResponseWriter, req *http.Request){
	res.Write([]byte(`Here are the saved games`))
}

func SaveGame(res http.ResponseWriter, req *http.Request){
	data := make(map[string]interface{})
	
	json.NewDecoder(req.Body).Decode(&data)
	fmt.Println(data)

    res.Header().Set("Content-Type", "application/json") 
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

func AddUserToLocation(res http.ResponseWriter, req *http.Request){
	data := make(map[string]interface{})
	
	json.NewDecoder(req.Body).Decode(&data)
	fmt.Println(data)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(data)
}

func RemoveUserFromLocation(res http.ResponseWriter, req *http.Request){
	data := make(map[string]interface{})
	
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil{
		utilities.SendJSON(http.StatusBadRequest, res, err.Error())
		return
	}
	
	fmt.Println(data)

	utilities.SendJSON(http.StatusOK, res, data)
}
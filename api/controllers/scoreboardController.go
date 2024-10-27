package controllers

import (
	"net/http"
	"DigitalScoreBoard/api/database"
)

// type ScoreBoardController struct {
// }

// func(s *ScoreBoardController) Init(){

// }

func GetAllUsersByLocation(res http.ResponseWriter, req *http.Request) {
	database.GetDB().Database("Adapt")
}

func SaveGame(res http.ResponseWriter, req *http.Request){

}

func AddUserToLocation(res http.ResponseWriter, req *http.Request){

}

func RemoveUserFromLocation(res http.ResponseWriter, req *http.Request){

}
package state

import(
	"DigitalScoreBoard/models"
)

var Users []models.User

func InitState(){
	Users = []models.User{{
		Username: "Darien",
		Points: 1200,
	},
	{
		Username: "Vicky",
		Points: 400,
	}, 
	{
		Username: "Richard",
		Points: 750,
	}, 
	{
		Username: "Kash",
		Points: 300,
	}}

}
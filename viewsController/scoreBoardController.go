package viewsController

import (
	"html/template"
	"net/http"
	"DigitalScoreBoard/models"
)

type ScoreBoardController struct{
	tmpl *template.Template
	users []models.User
}

func (s *ScoreBoardController) Init(){
	s.tmpl = template.Must(template.ParseGlob("templates/*"))
	s.users = []models.User{{
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

func (s *ScoreBoardController) HomePage(res http.ResponseWriter, req *http.Request){
	s.tmpl.ExecuteTemplate(res, "scoreboard.html", s.users)
}

func (s *ScoreBoardController) AddNewUserToGame(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil{
		http.Error(res, "Invalid form data", http.StatusBadRequest)
		return
	}

	newUser := req.FormValue("addUser")

	if newUser == "" {
		http.Error(res, "User cannot be empty.", http.StatusBadRequest)
		return
	}

	s.users = append(s.users, models.User{Username: newUser, Points: 0})

	//Execute a block element in the html file to add a new user card.
	s.tmpl.ExecuteTemplate(res, "user-card", models.User{Username: newUser, Points: 0})
}

func (s *ScoreBoardController) AddPointsToUser(res http.ResponseWriter, req *http.Request){

}

func (s *ScoreBoardController) SubtractPointsFromUser(res http.ResponseWriter, req *http.Request){

}

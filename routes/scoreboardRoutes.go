package routes

import (
	"github.com/go-chi/chi/v5"
	
	"DigitalScoreBoard/viewsController"
)

type ScoreBoardRoutes struct{
	Router     *chi.Mux
	controller viewsController.ScoreBoardController
}


func (s *ScoreBoardRoutes) Init(){
	s.Router = chi.NewRouter()
	
	s.controller.Init()
	s.Router.Get("/", s.controller.HomePage)
	s.Router.Post("/", s.controller.AddNewUserToGame)
	s.Router.Post("/addPoints/{username}", s.controller.AddPointsToUser)
	s.Router.Post("/subtractPoints/{username}", s.controller.SubtractPointsFromUser)
}

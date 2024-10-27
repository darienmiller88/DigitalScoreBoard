package routes

import (
	"github.com/go-chi/chi/v5"

	"DigitalScoreBoard/api/controllers"
)

type ScoreBoardRoutes struct{
	Router     *chi.Mux
	Controller controllers.ScoreBoardController
}


func (s *ScoreBoardRoutes) Init(){
	s.Router = chi.NewRouter()
	
}

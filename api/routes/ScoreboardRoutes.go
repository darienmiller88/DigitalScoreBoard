package routes

import (
	"github.com/go-chi/chi/v5"

	"DigitalScoreBoard/api/controllers"
)

type ScoreBoardRoutes struct{
	Router *chi.Mux
}

func (s *ScoreBoardRoutes) Init(){
	s.Router = chi.NewRouter()

	s.Router.Get("/get-location/{location-name}", controllers.GetLocation)
	s.Router.Get("/get-saved-games/{location-name}", controllers.GetAllSavedGames)
	s.Router.Get("/get-all-users/{location-name}", controllers.GetAllUsersByLocation)
	s.Router.Get("/get-all-locations", controllers.GetAllLocations)
	s.Router.Get("/get-saved-games", controllers.GetAllSavedGames)
	// s.Router.Get("/get-all-users", controllers.GetAllUsers)
	// s.Router.Post("/add-location", controllers.AddLocation)
	s.Router.Post("/save-game", controllers.SaveGame)
	s.Router.Post("/add-user-to-location/{location-name}", controllers.AddUserToLocation)
	s.Router.Delete("/remove-user-from-location/{location-name}/{username}", controllers.RemoveUserFromLocation)
}
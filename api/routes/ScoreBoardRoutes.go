package routes

import (
	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/httprate"

	"DigitalScoreBoard/api/controllers"
)

type ScoreBoardRoutes struct{
	Router *chi.Mux
}

func (s *ScoreBoardRoutes) Init(){
	s.Router = chi.NewRouter()

	//GET routes
	s.Router.Get("/get-location/{location-name}", controllers.GetLocation)
	s.Router.Get("/get-all-users/{location-name}", controllers.GetAllUsersByLocation)
	s.Router.Get("/get-saved-games/{location-name}", controllers.GetAllSavedGamesFromLocation)
	s.Router.Get("/get-current-team-games", controllers.GetAllSavedGames)
	s.Router.Get("/get-all-locations", controllers.GetAllLocations)
	s.Router.Get("/get-saved-games",  controllers.GetAllSavedGames)
	s.Router.Get("/get-all-users", controllers.GetAllUsers)
	
	//In order to rate limit POST, PUT, and DELETE routes, group them together away from the get routes, which allows
	//to attach a rate limiter middleware
	s.Router.Group(func(r chi.Router) {

		//POST routes
		r.Post("/save-game", controllers.SaveGame)
		r.Post("/add-user-to-location/{location-name}", controllers.AddUserToLocation)
		// s.Router.Post("/add-location", controllers.AddLocation)

		//PUT route(s)
		r.Put("/change-player-name/{location-name}", controllers.UpdatePlayerName)
		
		//DELETE route(s)
		r.Delete("/remove-user-from-location/{location-name}", controllers.RemoveUserFromLocation)
	})

}

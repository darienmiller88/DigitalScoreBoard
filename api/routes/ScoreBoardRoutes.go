package routes

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"

	"DigitalScoreBoard/api/controllers"
)

type ScoreBoardRoutes struct{
	Router *chi.Mux
}

func (s *ScoreBoardRoutes) Init(){
	s.Router = chi.NewRouter()

	//Group together GET routes to apply more lenient rate limiting to prevent server overload.
	s.Router.Group(func(r chi.Router) {
		r.Use(httprate.LimitByIP(100, 1 * time.Minute))

		//GET routes
		r.Get("/get-location/{location-name}", controllers.GetLocation)
		r.Get("/get-all-users/{location-name}", controllers.GetAllUsersByLocation)
		r.Get("/get-saved-games/{location-name}", controllers.GetAllSavedGamesFromLocation)
		r.Get("/get-current-team-games", controllers.GetAllSavedGames)
		r.Get("/get-all-locations", controllers.GetAllLocations)
		r.Get("/get-saved-games",  controllers.GetAllSavedGames)
		r.Get("/get-all-users", controllers.GetAllUsers)
	})
	
	//In order to rate limit POST, PUT, and DELETE routes, group them together away from the get routes, which allows
	//to attach a rate limiter middleware
	s.Router.Group(func(r chi.Router) {
		r.Use(httprate.LimitByIP(15, 1 * time.Minute))

		//POST routes
		r.Post("/save-game", controllers.SaveGame)
		r.Post("/add-user-to-location/{location-name}", controllers.AddUserToLocation)

		//PUT route(s)
		r.Put("/change-player-name/{location-name}", controllers.UpdatePlayerName)
		
		//DELETE route(s)
		r.Delete("/remove-user-from-location/{location-name}", controllers.RemoveUserFromLocation)
		r.Delete("/delete-save-game/{id}", controllers.DeleteSavedGame)
	})
}

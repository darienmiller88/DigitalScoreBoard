package routes

import "github.com/go-chi/chi/v5"

type UserRoutes struct {
	Router *chi.Mux
}

func (u *UserRoutes) Init(){
	u.Router = chi.NewRouter()
	// 
}
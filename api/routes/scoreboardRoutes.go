package routes

import (
	"github.com/go-chi/chi/v5"
	
)

type scoreBoardRoutes struct{
	Router     *chi.Mux
}


func (s *scoreBoardRoutes) Init(){
	s.Router = chi.NewRouter()
	
}

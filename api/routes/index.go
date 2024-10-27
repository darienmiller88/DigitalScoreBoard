package routes

import(
	"github.com/go-chi/chi/v5"
)

type Index struct{
	Router *chi.Mux
	sbRoutes ScoreBoardRoutes
}

func (i *Index) Init(){
	i.Router = chi.NewRouter()

	i.sbRoutes.Init()
	i.Router.Mount("/api/v1", i.sbRoutes.Router)
}
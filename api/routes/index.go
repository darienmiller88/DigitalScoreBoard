package routes

import(
	"github.com/go-chi/chi/v5"
)

type Index struct{
	Router *chi.Mux
	sbRoutes ScoreBoardRoutes
	userRoutes UserRoutes
}

func (i *Index) Init(){
	i.Router = chi.NewRouter()

	i.sbRoutes.Init()
	i.userRoutes.Init()

	i.Router.Mount("/scoreboard", i.sbRoutes.Router)
	i.Router.Mount("/users", i.userRoutes.Router)
}
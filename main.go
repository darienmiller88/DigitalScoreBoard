package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5/middleware"

	"DigitalScoreBoard/routes"
)

type User struct{
	Username string
	Points int
}

func main(){
	godotenv.Load()

	router := chi.NewRouter()
	scoreBoardRoutes := routes.ScoreBoardRoutes{}	

	scoreBoardRoutes.Init()
	router.Use(middleware.Logger, middleware.RealIP, middleware.Recoverer)
    router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.Mount("/", scoreBoardRoutes.Router)

	fmt.Println("listening on port:", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
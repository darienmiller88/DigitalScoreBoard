package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/routes"
)

func main(){
	godotenv.Load()

	router := chi.NewRouter()
	index := routes.Index{}
	newCors := cors.New(cors.Options{
		AllowedOrigins:   []string{ "http://localhost:5173", "http://127.0.0.1:5173", "https://adaptscoreboard.netlify.app" },
		AllowedMethods:   []string{ "GET", "POST", "PUT", "DELETE", "OPTIONS" },
		AllowCredentials: true,
	})

	router.Use(middleware.Logger, middleware.RealIP, middleware.Recoverer, newCors.Handler)
	
	//Initialize both the subrouter, as well as the mongoDB database instance.
	index.Init()
	database.Init()

	defer database.DisconnectClient()

	router.Mount("/api/v1", index.Router)
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`Welcome to my Adapt scoreboard API!`))
	})

	fmt.Println("listening on port:", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
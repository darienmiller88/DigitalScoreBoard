package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5/middleware"
	"DigitalScoreBoard/api/database"
	"DigitalScoreBoard/api/routes"
)

func main(){
	godotenv.Load()

	index := routes.Index{}
	router := chi.NewRouter()

	//Initialize both the subrouter, as well as the mongoDB database instance.
	index.Init()
	database.Init()

	router.Use(middleware.Logger, middleware.RealIP, middleware.Recoverer)
	
	router.Mount("/", index.Router)
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`Hello world!`))
	})

	fmt.Println("listening on port:", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
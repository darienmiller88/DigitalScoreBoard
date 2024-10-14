package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	godotenv.Load()

	router := chi.NewRouter()

	router.Use(middleware.Logger, middleware.RealIP, middleware.Recoverer)

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`Hello world blah!`))
	})
   
	fmt.Println("listening on port:", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
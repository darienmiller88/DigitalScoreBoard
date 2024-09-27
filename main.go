package main

import(
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	router := chi.NewRouter()

	router.Get("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte(`Hello world`))
	})

	fmt.Println("listening on port:")
	http.ListenAndServe(":8080", router)
}
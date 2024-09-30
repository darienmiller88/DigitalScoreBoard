package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	router := chi.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	tmpl := template.Must(template.ParseGlob("./templates/*"))

    router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Get("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`Hello world`))
	})

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		tmpl.ExecuteTemplate(res, "scoreboard.html", map[string][]string{
			"Users": {"Darien", "Vicky", "Richard", "Kash"},
		})
	})

	fmt.Println("listening on port:", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
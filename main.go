package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	godotenv.Load()

	router := chi.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	tmpl := template.Must(template.ParseGlob("templates/*"))
	users := []string{"Darien", "Vicky", "Richard", "Kash"}

	router.Use(middleware.Logger, middleware.RealIP, middleware.Recoverer)
    router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Get("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`Hello world`))
	})

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		tmpl.ExecuteTemplate(res, "scoreboard.html", map[string][]string{
			"Users": users,
		})
	})

	router.Post("/", func(res http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil{
			http.Error(res, "Invalid form data", http.StatusBadRequest)
            return
		}

		newUser := req.FormValue("addUser")

		if newUser == "" {
			http.Error(res, "User cannot be empty.", http.StatusBadRequest)
			return
		}

		users = append(users, newUser)

		//Execute a block element in the html file to add a new user to the "users" ul element.
		tmpl.ExecuteTemplate(res, "user-element", newUser)
	})

	fmt.Println("listening on port:", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)
}
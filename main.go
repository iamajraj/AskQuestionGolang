package main

import (
	"ask-anon-ques/controllers"
	"ask-anon-ques/db"
	"ask-anon-ques/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main(){
	db.ConnectDB();

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.SendJSON(utils.JSONT{
			"hello":"world",
		},&w)
	})

	router.Route("/users", func(r chi.Router) {
		r.Get("/", controllers.GetUsers)
		r.Post("/", controllers.CreateUser)
		r.Post("/{id}", controllers.CreateQuestion)
	})

	fmt.Println("Server started... http://127.0.0.1:4000")
	http.ListenAndServe("127.0.0.1:4000", router)
}
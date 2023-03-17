package routes

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/today", controllers.Today)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/complete/{id}", controllers.Complete).Methods("POST")

	return route
}

package routes

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.NotPostponed)
	route.HandleFunc("/postponed", controllers.Postponed)
	route.HandleFunc("/not-postponed", controllers.NotPostponedNew)
	route.HandleFunc("/todo-incomplete", controllers.NotCompleted)
	route.HandleFunc("/todo-completed", controllers.Completed)
	route.HandleFunc("/all", controllers.All)
	route.HandleFunc("/focus", controllers.Focus)
	route.HandleFunc("/ping", controllers.Ping)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/complete/{id}", controllers.Complete).Methods("POST")

	return route
}

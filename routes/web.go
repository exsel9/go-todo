package routes

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.NotDeferred)
	route.HandleFunc("/not-completed", controllers.NotCompleted)
	route.HandleFunc("/all", controllers.All)
	route.HandleFunc("/focus", controllers.Focus)
	route.HandleFunc("/ping", controllers.Ping)
	route.HandleFunc("/add", controllers.Add).Methods("POST")

	return route
}

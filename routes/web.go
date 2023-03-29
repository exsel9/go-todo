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
	route.HandleFunc("/incomplete/{id}", controllers.Incomplete).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	route.HandleFunc("/focus/{id}", controllers.Focus).Methods("POST")
	route.HandleFunc("/unfocused/{id}", controllers.Unfocused).Methods("POST")
	route.HandleFunc("/postpone/{id}", controllers.Postpone).Methods("POST")

	return route
}

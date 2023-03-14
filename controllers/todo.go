package controllers

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/models"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"net/http"
)

var (
	id        int
	item      string
	completed int
	view      = template.Must(template.ParseFiles("./views/index.html"))
	database  = config.Database()
)

func Show(w http.ResponseWriter, _ *http.Request) {
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		log.Error(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &item, &completed)

		if err != nil {
			log.Error(err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	item := r.FormValue("item")

	_, err := database.Exec(`INSERT INTO todos (item) VALUE (?)`, item)

	if err != nil {
		log.Error(err)
	}

	http.Redirect(w, r, "/", 302)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}

	http.Redirect(w, r, "/", 302)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}

	http.Redirect(w, r, "/", 302)
}

func UnComplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE todos SET completed = 0 WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}

	http.Redirect(w, r, "/", 302)
}

func Healthz(w http.ResponseWriter, _ *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Error(err)
	}
}

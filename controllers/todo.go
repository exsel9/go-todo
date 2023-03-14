package controllers

import (
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/item_dao"
	"github.com/ichtrojan/go-todo/models"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"net/http"
)

var (
	view     = template.Must(template.ParseFiles("./views/index.html"))
	database = config.Database()
	itemDAO  = item_dao.New(database)
)

func Show(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.All()

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	it := r.FormValue("item")
	log.WithFields(log.Fields{"description": it}).Info("Add new TodoItem. Saving to database.")

	itemDAO.Add(models.Todo{
		Item: it,
	})

	http.Redirect(w, r, "/", 302)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemDAO.Delete(vars["id"])

	http.Redirect(w, r, "/", 302)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemDAO.MarkAsComplete(vars["id"])

	http.Redirect(w, r, "/", 302)
}

func UnComplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemDAO.MarkAsUnComplete(vars["id"])

	http.Redirect(w, r, "/", 302)
}

func Ping(w http.ResponseWriter, _ *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Error(err)
	}
}

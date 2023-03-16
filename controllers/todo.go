package controllers

import (
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/item_dao"
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

func NotCompleted(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.NotCompleted()

	_ = view.Execute(w, todos)
}

func Postponed(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.Postponed()

	_ = view.Execute(w, todos)
}

func NotPostponed(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.NotPostponed()

	_ = view.Execute(w, todos)
}

func All(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.All()

	_ = view.Execute(w, todos)
}

func Focus(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.Focus()

	_ = view.Execute(w, todos)
}

func Add(w http.ResponseWriter, r *http.Request) {

	it := r.FormValue("item")
	log.WithFields(log.Fields{"description": it}).Info("Add new TodoItem. Saving to database.")

	itemDAO.Add(it)

	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"status": "success"}`)
	if err != nil {
		log.Error(err)
	}
}

func Ping(w http.ResponseWriter, _ *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		log.Error(err)
	}
}

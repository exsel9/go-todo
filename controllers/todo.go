package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		log.Error(err)
	}
}

func Completed(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.Completed()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		log.Error(err)
	}
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	itemDAO.MarkAsComplete(id)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err := io.WriteString(w, `{"status": "success"}`)
	if err != nil {
		log.Error(err)
	}
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
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}

	it := r.Form["item"][0]
	log.WithFields(log.Fields{"item": it}).Info("Add new TodoItem. Saving to database.")

	id := itemDAO.Add(it)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp := fmt.Sprintf(`{"status": "success", "id": %d}`, id)
	_, err = io.WriteString(w, resp)
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

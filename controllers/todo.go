package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/item_dao"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var (
	database = config.Database()
	itemDAO  = item_dao.New(database)
)

func Today(w http.ResponseWriter, _ *http.Request) {
	todos := itemDAO.Today()

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

func Incomplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	itemDAO.MarkAsIncomplete(id)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err := io.WriteString(w, `{"status": "success"}`)
	if err != nil {
		log.Error(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	itemDAO.Delete(id)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err := io.WriteString(w, `{"status": "success"}`)
	if err != nil {
		log.Error(err)
	}
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

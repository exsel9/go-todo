package main

import (
	"github.com/ichtrojan/go-todo/routes"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}

	log.Info("Starting Todolist API server")

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}

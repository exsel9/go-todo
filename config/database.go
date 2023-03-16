package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"os"
)

func Database() *sql.DB {
	database, err := sql.Open("mysql", credentials())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE IF NOT EXISTS gotodo`)

	if err != nil {
		log.Error(err)
	}

	_, err = database.Exec(`USE gotodo`)

	if err != nil {
		log.Error(err)
	}

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    focused BOOLEAN DEFAULT FALSE,
		    repeated BOOLEAN DEFAULT FALSE,
		    postponed_until_date DATE NOT NULL DEFAULT (CURRENT_DATE),
		    PRIMARY KEY (id)
		);
	`)
	if err != nil {
		log.Error(err)
	}

	return database
}

func credentials() string {
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("DB_USER not set in .env")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		log.Fatal("DB_PASS not set in .env")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("DB_HOST not set in .env")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		log.Fatal("DB_PORT not set in .env")
	}

	return fmt.Sprintf("%s:%s@(%s:%s)/?charset=utf8&parseTime=True", user, pass, host, port)
}

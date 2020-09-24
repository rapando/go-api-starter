package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func DbConnect() (db *sql.DB){
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Println("Unable to run query in db because ", err)
		os.Exit(3)
	}

	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Second)
	log.Println("Db connection successful")
	return
}

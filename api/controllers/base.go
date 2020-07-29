package controllers

import (
	"api/api/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Db     *sql.DB
	Router *mux.Router
}

func (s *Server) dbConnect() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHost, dbName)
	if s.Db, err = sql.Open("mysql", dbURI); err != nil {
		log.Println("Unable to run query in db because ", err)
		os.Exit(3)
	}

	s.Db.SetMaxIdleConns(64)
	s.Db.SetMaxOpenConns(100)
	s.Db.SetConnMaxLifetime(10 * time.Second)
	log.Println("Db connection successful")
}

func (s *Server) Init() {
	utils.InitLogger()
	s.dbConnect()
	s.initRoutes()
}

func (s *Server) Run() {
	port := ":" + os.Getenv("PORT")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type", "content-length", "accept-encoding", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT"})

	log.Println("Listening on port ", port)

	if err := http.ListenAndServe(port, handlers.CORS(origins, headers, methods)(s.Router)); err != nil {
		log.Println("Unable to start app because ", err)
		s.Db.Close()
	} else {
		log.Println("Hey hey hey")
	}
}

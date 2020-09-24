package controllers

import (
	"api/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Db     *sql.DB
	Router *mux.Router
}

func (s *Server) Init() {
	utils.InitLogger()
	s.Db = utils.DbConnect()
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

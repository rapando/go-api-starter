package controllers

import (
	"github.com/gorilla/mux"
	"log"
)

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()
	log.Println("Initializing routes")

	s.Router.HandleFunc("/", s.Home).Methods("GET")
}

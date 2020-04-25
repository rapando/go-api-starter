package controllers

import "github.com/gorilla/mux"

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()
	s.Log("Initializing routes")

	s.Router.HandleFunc("/", s.Home).Methods("GET")
}

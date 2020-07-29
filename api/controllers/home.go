package controllers

import (
	"api/api/middleware"
	"log"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome home...")
	middleware.OkResponse(w, 200, "Ok")
}

package controllers

import (
	"api/api/middleware"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	s.Log("Welcome home...")
	middleware.OkResponse(w, 200, "Ok")
}

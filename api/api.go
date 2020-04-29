package api

import (
	"api/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	godotenv.Load()
	server.Init()
	server.Run()

}

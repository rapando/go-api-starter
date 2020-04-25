package api

import (
	"api/api/controllers"
	"os"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	godotenv.Load()
	if err := server.Init(); err != nil {
		server.Log("Unable to init server because ", err)
		os.Exit(3)
	}
	server.Run()

}

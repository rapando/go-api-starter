package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitLogger() {
	logFolder := os.Getenv("LOG_FOLDER")
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s%s.log", logFolder+"app-", "%Y-%m-%d"),
		rotatelogs.WithLinkName(logFolder+"link.log"),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(10000),
	)
	if err != nil {
		fmt.Println("Failed to initialize log file ", err.Error())
		log.SetOutput(os.Stderr)
		return
	}
	if os.Getenv("ENV") == "dev" {
		log.SetOutput(os.Stderr)
		return
	}
	log.SetOutput(writer)
	return
}
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
		fmt.Sprintf("%s%s.log", logFolder+"app-", "%Y-%m-%d.%H%M"),
		rotatelogs.WithLinkName(logFolder+"link.log"),
		rotatelogs.WithRotationTime(time.Hour*12),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(10000),
	)
	if err != nil {
		fmt.Println("Failed to initialize log file ", err.Error())
	}
	log.SetOutput(writer)
	return
}

func Log(msg ...interface{}) {
	env := os.Getenv("ENV")
	switch env {
	case "dev":
		fmt.Println(msg...)
	case "test":
		log.Println(msg...)
	default:
		msg = append(msg, "\n----------------------------")
		log.Println(msg...)
	}

}

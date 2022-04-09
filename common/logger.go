package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var env string

var fileLogger *log.Logger

func InitLogger() {
	env = viper.GetString("environment")

	timeStr := time.Now().Format("20060102-150405")

	if env == "production" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}
		logFilePath := wd + "/logs/"
		logFileName := timeStr + ".log"
		fileName := path.Join(logFilePath, logFileName)

		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			panic(err.Error())
		}
		f, err := os.Create(fileName)
		if err != nil {
			panic(err.Error())
		}

		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(f)
		fileLogger = log.New(f, "\r\n", log.LstdFlags)
		log.SetOutput(f)

		fmt.Println("==== logger env: production ====")
	}
}

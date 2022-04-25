package common

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
	"ustoj-master/scheduler/model"

	"github.com/sirupsen/logrus"
)

var env string

var fileLogger *log.Logger
var once sync.Once
var instance *logrus.Logger

var logger = LogInstance()

const (
	defaultLogFormat       = "[%lvl%] %time% | %msg%"
	defaultTimestampFormat = time.RFC3339
)

type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	if len(entry.Data) > 0 {
		fields, _ := json.Marshal(entry.Data)
		output = output + " | " + string(fields)
	}

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1) + "\n"

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}

func InitLogger(appConfig model.Config) {
	logger.Formatter = new(Formatter)
	// logger.SetLevel(config.Log.Level)

	if appConfig.Logger.WriteFile {
		// TODO: set log path from config
		// generate file name
		wd, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}
		timeStr := time.Now().Format("20060102-150405")
		logFilePath := wd + "/logs/"
		logFileName := timeStr + ".log"
		fileName := path.Join(logFilePath, logFileName)
		if err := os.Mkdir(logFilePath, 0755); err != nil {
			fmt.Println(err.Error())
		}
		if _, err := os.Stat(fileName); err != nil {
			if _, err := os.Create(fileName); err != nil {
				fmt.Println(err.Error())
			}
		}
		writeToFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("err", err)
		}

		writers := io.MultiWriter(os.Stdout, writeToFile)
		logger.Out = writers

		logger.Infoln("log file save in: " + fileName)
	}

}

func LogInstance() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()
	})
	return instance
}

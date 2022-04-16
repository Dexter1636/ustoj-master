package utils

import (
	"encoding/json"
	"log"
)

func LogBody(body interface{}, tag string) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("[%s] %s\n", tag, string(jsonBytes))
	}
}

func LogReqRespBody(req interface{}, resp interface{}, tag string) {
	LogBody(req, tag+".req")
	LogBody(resp, tag+".resp")
}

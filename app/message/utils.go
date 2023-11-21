package message

import (
	"encoding/json"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ParseOrderToMessage(request interface{}) ([]byte, error) {
	jsonObj, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	return jsonObj, err
}

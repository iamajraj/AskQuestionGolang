package utils

import (
	"encoding/json"
	"net/http"
)

type JSONT = map[string]interface{}

func SendJSON(data interface{}, w* http.ResponseWriter){
	json_marshal, err := json.Marshal(data)

	if err != nil {
		(*w).Write([]byte("Something Went Wrong"))
		return
	}

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(json_marshal)
}
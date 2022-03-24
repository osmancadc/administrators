package main

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, status int, message interface{}, isJson ...bool) {
	if len(isJson) > 0 && isJson[0] {
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			SendResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(jsonResponse)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(message)
	}
}

func GetCriticalityId(criticality string) int {
	if criticality == "low" {
		return 1
	} else if criticality == "medium" {
		return 2
	} else if criticality == "high" {
		return 3
	} else {
		return -1
	}
}

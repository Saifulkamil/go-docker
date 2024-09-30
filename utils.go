package main

import (
	"net/http"
	"encoding/json"
)

type ResponJSON struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

func sendResponse(w http.ResponseWriter, message string, data any, status... int) {
	responseStatus := http.StatusOK
    if len(status) > 0 {
        responseStatus = status[0]
    }
	res := ResponJSON{
		Message: message,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	json.NewEncoder(w).Encode(res)
}
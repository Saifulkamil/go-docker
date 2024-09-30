package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type ResponJSON struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func sendResponse(w http.ResponseWriter, message string, data any, status ...int) {
	responseStatus := http.StatusOK
	if len(status) > 0 {
		responseStatus = status[0]
	}
	res := ResponJSON{
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	json.NewEncoder(w).Encode(res)
}

func validateRequired(value any) bool {
	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return false
		}
	default:
		if reflect.ValueOf(value).IsZero() {
			return false
		}
	}
	return true
}

func validateExists(table string, id int) bool {
	query := fmt.Sprintf("SELECT id FROM %s WHERE id = ?", table)

	var retrievedID int
	err := db.QueryRow(query, id).Scan(&retrievedID)

	return err == nil
}

func validateNumeric(value any) bool {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case string:
		if _, err := strconv.Atoi(v); err == nil {
			return true
		}
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			return true
		}
	}
	return false
}

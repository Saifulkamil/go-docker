package main

import (
	"encoding/json"
	"net/http"
)

func getCategories(w http.ResponseWriter, r *http.Request) {

	type ResJSON struct {
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")
    message := ResJSON{Message: "Hello, from handler!"}
    json.NewEncoder(w).Encode(message)
}

// func createCategory(w http.ResponseWriter, r *http.Request) {}
// func getItems(w http.ResponseWriter, r *http.Request) {}
// func getItemByID(w http.ResponseWriter, r *http.Request) {}
// func createItem(w http.ResponseWriter, r *http.Request) {}
// func updateItem(w http.ResponseWriter, r *http.Request) {}
// func deleteItem(w http.ResponseWriter, r *http.Request) {}
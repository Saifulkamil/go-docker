package main

import (
	"encoding/json"
	// "log"
	"net/http"
)


func getCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var categories []Category
	
	for rows.Next() {
		var cat Category
		rows.Scan(&cat.ID, &cat.Name)
		categories = append(categories, cat)
	}

	type ResponJSON struct {
		Message string `json:"message"`
		Data []Category `json:"data"`
	}

	res := ResponJSON{
		Message: "Categories loaded successfully",
		Data: categories,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// func createCategory(w http.ResponseWriter, r *http.Request) {}
func getItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var items []Item
	
	for rows.Next() {
		var item Item
		rows.Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.CreatedAt)
		items = append(items, item)
	}

	type ResponJSON struct {
		Message string `json:"message"`
		Data []Item `json:"data"`
	}

	res := ResponJSON{
		Message: "Items loaded successfully",
		Data: items,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
// func getItemByID(w http.ResponseWriter, r *http.Request) {}
// func createItem(w http.ResponseWriter, r *http.Request) {}
// func updateItem(w http.ResponseWriter, r *http.Request) {}
// func deleteItem(w http.ResponseWriter, r *http.Request) {}
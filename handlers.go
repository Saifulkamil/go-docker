package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"database/sql"
)

// Category Handlers
func categoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET" :
			getCategories(w)
		case "POST" :
			createCategory(w, r)
		default :
			sendResponse(w, "Method not allowed", nil, http.StatusMethodNotAllowed)
			return
	}
}

func getCategories(w http.ResponseWriter) {
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

	sendResponse(w, "Categories loaded successfully", categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}
	var params Params
	json.NewDecoder(r.Body).Decode(&params)

	errs := make(map[string]string)

	if (!validateRequired(params.Name)) {
		errs["name"] = "name is required"
	}

	if len(errs) > 0 {
		sendResponse(w, "Invalid field", errs, http.StatusBadRequest)
		return
	}

	sql := "INSERT INTO categories (name) VALUES (?)"
	_, err := db.Exec(sql, params.Name)
	if err != nil {
		sendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	
	sendResponse(w, "Category created successfully", nil)
}
// End Category Handlers

// Item Handler
func itemHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	log.Println()
	if len(parts) > 1 && len(parts) <= 3 {

		switch r.Method {
			case "GET" :
				if parts[2] != "" {
					if (!validateNumeric(parts[2])) {
						sendResponse(w, "ID should be numeric", nil, http.StatusBadRequest)
						return
					}
					id, _ := strconv.Atoi(parts[2])
					getItemByID(w, id)
					return
				}
				getItems(w)
				return
			case "POST" :
				createItem(w, r)
				return
			case "PUT" :
				if parts[2] != "" {
					if (!validateNumeric(parts[2])) {
						sendResponse(w, "ID should be numeric", nil, http.StatusBadRequest)
						return
					}
					id, _ := strconv.Atoi(parts[2])
					updateItem(w, r, id)
					return
				}
			case "DELETE" :
				if parts[2] != "" {
					if (!validateNumeric(parts[2])) {
						sendResponse(w, "ID should be numeric", nil, http.StatusBadRequest)
						return
					}
					id, _ := strconv.Atoi(parts[2])
					deleteItem(w, id)
					return
				}
		}
		sendResponse(w, "Method not allowed", nil, http.StatusMethodNotAllowed)
		return
	}
	sendResponse(w, "Route not found", nil, http.StatusNotFound)
}

func getItems(w http.ResponseWriter) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		sendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var items []Item
	
	for rows.Next() {
		var item Item
		rows.Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.CreatedAt)
		items = append(items, item)
	}

	sendResponse(w, "Items loaded successfully", items)
}

type ParamsItem struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CategoryID int `json:"category_id"`
}
func itemValidator(params ParamsItem) (map[string]string){
	errs := make(map[string]string)
	
	if (!validateRequired(params.Name)) {
		errs["name"] = "name is required"
	}
	if (!validateRequired(params.Description)) {
		errs["description"] = "description is required"
	}
	if (!validateRequired(params.Price)) {
		errs["price"] = "price is required"
	}
	if (!validateNumeric(params.Price)) {
		errs["price"] = "price should be numeric"
	}
	if (!validateRequired(params.CategoryID)) {
		errs["category_at"] = "category_at is required"
	}
	if (!validateExists("categories", params.CategoryID)) {
		errs["category_at"] = "category_at does not exist"
	}
	if (!validateNumeric(params.CategoryID)) {
		errs["category_at"] = "caategory_id should be numeric"
	}
	return errs
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var params ParamsItem
	json.NewDecoder(r.Body).Decode(&params)

	errs := itemValidator(params)
	if len(errs) > 0 {
		sendResponse(w, "Invalid field", errs, http.StatusBadRequest)
		return
	}

	sql := "INSERT INTO items (name, category_id, description, price) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(sql, params.Name, params.CategoryID, params.Description, params.Price)
	if err != nil {
		sendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	sendResponse(w, "Item created successfully", nil)
}

func getItemByID(w http.ResponseWriter, id int) {
	var item Item
	query := "SELECT id, category_id, name, description, price, created_at FROM items WHERE id=?"
	err := db.QueryRow(query, id).Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			sendResponse(w, "Item not found", nil, http.StatusNotFound)
			return
		}
		sendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	sendResponse(w, "Item loaded successfully", item)
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var params ParamsItem
	json.NewDecoder(r.Body).Decode(&params)

	errs := itemValidator(params)
	if len(errs) > 0 {
		sendResponse(w, "Invalid field", errs, http.StatusBadRequest)
		return
	}

	sql := "UPDATE items SET name=?, category_id=?, description=?, price=? WHERE id=?"
	_, err := db.Exec(sql, params.Name, params.CategoryID, params.Description, params.Price, id)
	if err != nil {
		sendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	sendResponse(w, "Item updated successfully", nil)
}

func deleteItem(w http.ResponseWriter, id int) {
	sql := "DELETE FROM items WHERE id=?"
	_, err := db.Exec(sql, id)
	if err != nil {
		sendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	sendResponse(w, "Item deleted successfully", nil)

}
// End Item Handler
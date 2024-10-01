package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pari_test/utils"
	"strconv"
	"strings"
)

// Category Handlers
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCategories(w)
	case "POST":
		createCategory(w, r)
	default:
		utils.SendResponse(w, "Method not allowed", nil, http.StatusMethodNotAllowed)
		return
	}
}

func getCategories(w http.ResponseWriter) {
	rows, err := utils.DB.Query("SELECT * FROM categories")
	if err != nil {
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var cat Category
		rows.Scan(&cat.ID, &cat.Name)
		categories = append(categories, cat)
	}

	utils.SendResponse(w, "Categories loaded successfully", categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}
	var params Params
	json.NewDecoder(r.Body).Decode(&params)

	errs := make(map[string]string)

	if !utils.ValidateRequired(params.Name) {
		errs["name"] = "name is required"
	}

	if len(errs) > 0 {
		utils.SendResponse(w, "Invalid field", errs, http.StatusBadRequest)
		return
	}

	sql := "INSERT INTO categories (name) VALUES (?)"
	_, err := utils.DB.Exec(sql, params.Name)
	if err != nil {
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, "Category created successfully", nil)
}

// End Category Handlers

// Item Handler
func ItemHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) <= 3 {
		switch r.Method {
		case "GET":
			if len(parts) > 2 && parts[2] != "" {
				if !utils.ValidateNumeric(parts[2]) {
					utils.SendResponse(w, "ID should be numeric", nil, http.StatusBadRequest)
					return
				}
				id, _ := strconv.Atoi(parts[2])
				getItemByID(w, id)
				return
			}
			getItems(w, r)
			return
		case "POST":
			createItem(w, r)
			return
		case "PUT":
			if len(parts) > 2 && parts[2] != "" {
				if !utils.ValidateNumeric(parts[2]) {
					utils.SendResponse(w, "ID should be numeric", nil, http.StatusBadRequest)
					return
				}
				id, _ := strconv.Atoi(parts[2])
				updateItem(w, r, id)
				return
			}
		case "DELETE":
			if len(parts) > 2 && parts[2] != "" {
				if !utils.ValidateNumeric(parts[2]) {
					utils.SendResponse(w, "ID should be numeric", nil, http.StatusBadRequest)
					return
				}
				id, _ := strconv.Atoi(parts[2])
				deleteItem(w, id)
				return
			}
		}
		utils.SendResponse(w, "Method not allowed", nil, http.StatusMethodNotAllowed)
		return
	}
	utils.SendResponse(w, "Route not found", nil, http.StatusNotFound)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	sortBy := r.URL.Query().Get("sort")
	sortOrder := r.URL.Query().Get("order")

	query := "SELECT items.id, items.category_id, items.name, items.description, items.price, items.created_at FROM items LEFT JOIN categories ON categories.id=items.category_id"

	if search != "" {
		query += " WHERE items.name LIKE '%" + search + "%' OR categories.name LIKE '%" + search + "%'"
	}

	if strings.ToLower(sortBy) == "name" || strings.ToLower(sortBy) == "price" {
		if strings.ToLower(sortOrder) == "desc" {
			sortOrder = "DESC"
		} else {
			sortOrder = "ASC"
		}
		query += " ORDER BY items." + sortBy + " " + sortOrder
	}

	rows, err := utils.DB.Query(query)
	if err != nil {
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item
		rows.Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.CreatedAt)
		items = append(items, item)
	}

	utils.SendResponse(w, "Items loaded successfully", items)
}

type ParamsItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
}

func itemValidator(params ParamsItem) map[string]string {
	errs := make(map[string]string)

	if !utils.ValidateRequired(params.Name) {
		errs["name"] = "name is required"
	}
	if !utils.ValidateRequired(params.Description) {
		errs["description"] = "description is required"
	}
	if !utils.ValidateRequired(params.Price) {
		errs["price"] = "price is required"
	}
	if !utils.ValidateNumeric(params.Price) {
		errs["price"] = "price should be numeric"
	}
	if !utils.ValidateRequired(params.CategoryID) {
		errs["category_at"] = "category_at is required"
	}
	if !utils.ValidateExists("categories", params.CategoryID) {
		errs["category_at"] = "category_at does not exist"
	}
	if !utils.ValidateNumeric(params.CategoryID) {
		errs["category_at"] = "caategory_id should be numeric"
	}
	return errs
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var params ParamsItem
	json.NewDecoder(r.Body).Decode(&params)

	errs := itemValidator(params)
	if len(errs) > 0 {
		utils.SendResponse(w, "Invalid field", errs, http.StatusBadRequest)
		return
	}

	sql := "INSERT INTO items (name, category_id, description, price) VALUES (?, ?, ?, ?)"
	_, err := utils.DB.Exec(sql, params.Name, params.CategoryID, params.Description, params.Price)
	if err != nil {
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, "Item created successfully", nil)
}

func getItemByID(w http.ResponseWriter, id int) {
	var item Item
	query := "SELECT id, category_id, name, description, price, created_at FROM items WHERE id=?"
	err := utils.DB.QueryRow(query, id).Scan(&item.ID, &item.CategoryID, &item.Name, &item.Description, &item.Price, &item.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendResponse(w, "Item not found", nil, http.StatusNotFound)
			return
		}
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, "Item loaded successfully", item)
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var params ParamsItem
	json.NewDecoder(r.Body).Decode(&params)

	if !utils.ValidateExists("items", id) {
		utils.SendResponse(w, "Item not found", nil, http.StatusNotFound)
		return
	}

	errs := itemValidator(params)
	if len(errs) > 0 {
		utils.SendResponse(w, "Invalid field", errs, http.StatusBadRequest)
		return
	}

	sql := "UPDATE items SET name=?, category_id=?, description=?, price=? WHERE id=?"
	_, err := utils.DB.Exec(sql, params.Name, params.CategoryID, params.Description, params.Price, id)
	if err != nil {
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, "Item updated successfully", nil)
}

func deleteItem(w http.ResponseWriter, id int) {

	if !utils.ValidateExists("items", id) {
		utils.SendResponse(w, "Item not found", nil, http.StatusNotFound)
		return
	}

	sql := "DELETE FROM items WHERE id=?"
	_, err := utils.DB.Exec(sql, id)
	if err != nil {
		utils.SendResponse(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, "Item deleted successfully", nil)
}

// End Item Handler

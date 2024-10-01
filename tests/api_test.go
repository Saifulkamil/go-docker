package tests

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	// "strings"
	"pari_test/app"
	"pari_test/utils"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Helper functions
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
		return nil
	}

	err = utils.DBConnection()
	if err != nil {
		log.Println(err)
		return nil
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.URL.Path)
		paths := strings.Split(r.URL.Path, "/")
		switch paths[1] {
		case "categories":
			app.CategoryHandler(w, r)
		case "items":
			app.ItemHandler(w, r)
		}
	})
	handler.ServeHTTP(rr, req)
	return rr
}

func TestGetCategories(t *testing.T) {
	req, _ := http.NewRequest("GET", "/categories", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateCategory(t *testing.T) {
	payload := []byte(`{"name": "New Category"}`)
	req, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	// Test for invalid data
	payload = []byte(`{"name": ""}`)
	req, _ = http.NewRequest("POST", "/categories", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestGetItems(t *testing.T) {
	req, _ := http.NewRequest("GET", "/items", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateItem(t *testing.T) {
	payload := []byte(`{"name": "New Item", "description": "Item Description", "price": 10.5, "category_id": 1}`)
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	// Test for invalid data
	payload = []byte(`{"name": "", "description": "No Name", "price": "not_numeric", "category_id": "wrong_type"}`)
	req, _ = http.NewRequest("POST", "/items", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestUpdateItem(t *testing.T) {
	payload := []byte(`{"name": "Updated Item", "description": "Updated Description", "price": 15.5, "category_id": 1}`)
	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	// Test for invalid ID
	req, _ = http.NewRequest("PUT", "/items/invalid_id", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestDeleteItem(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	// Test for invalid ID
	req, _ = http.NewRequest("DELETE", "/items/invalid_id", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

// Helper function to check response code
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

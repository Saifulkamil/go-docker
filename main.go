package main

import (
	"log"
	"net/http"
	"os"
	"pari_test/app"
	"pari_test/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

	// Connect to DB
	db_name := os.Getenv("DB_NAME")
	err = utils.DBConnection(db_name)
	if err != nil {
		log.Println(err)
	}else {
		log.Println("Database connected succesfully!")
	}

	// Register Routes
	RegisterRoutes()

	// Start server at port 8080
	log.Println("Server listen at port 8080")
	http.ListenAndServe("localhost:8080", nil)
}

func RegisterRoutes() {

	// Categories
	http.HandleFunc("/categories", app.CategoryHandler)

	// Items
	http.HandleFunc("/items/", app.ItemHandler)
}
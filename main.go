package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

	// Connect to DB
	db, err = dbConnection()
	if err != nil {
		log.Println(err)
	}else {
		log.Println("Database connected succesfully!")
	}

	// Register Routes
	registerRoutes()

	// Start server at port 8080
	log.Println("Server listen at port 8080")
	http.ListenAndServe("localhost:8080", nil)
}

func registerRoutes() {

	// Categories
	http.HandleFunc("/categories", categoryHandler)

	// Items
	http.HandleFunc("/items/", itemHandler)
}

func dbConnection() (*sql.DB, error) {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass, db_host, db_port, db_name)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	// defer db.Close()
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

	// Connect to DB
	var err error
	db, err = dbConnection()
	if err != nil {
		log.Fatal(err)
	}else {
		log.Println("Database connected succesfully!")
	}

	// Register Routes
	registerRoutes()

	// Start server at port 8080
	http.ListenAndServe(":8080", nil)
	log.Println("Server listen at port 8080")
}

func registerRoutes() {

	// Categories
	http.HandleFunc("/categories", getCategories)
}

func dbConnection() (*sql.DB, error) {
	db_user := "root"
	db_pass := ""
	db_name := "db_pari_test"
	db, err := sql.Open("mysql", db_user + ":" + db_pass + "@tcp(127.0.0.1:3306)/" + db_name)

	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
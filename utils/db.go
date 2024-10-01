package utils

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func DBConnection(db_name string) (error) {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass, db_host, db_port, db_name)

	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}
	
	err = DB.Ping()
	return err
}
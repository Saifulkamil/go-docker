package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/categories", getCategories)
	log.Println("Server listen at port 8080")
	http.ListenAndServe(":8080", nil)
}
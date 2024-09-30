package main

import (
	"time"
)

type Category struct {
	ID 	 int	`json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID 	 int	`json:"id"`
	CategoryID string `json:"category_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
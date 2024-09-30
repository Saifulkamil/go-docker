package main

type Category struct {
	ID 	 int	`json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID 	 int	`json:"id"`
	CategoryID int `json:"category_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CreatedAt string `json:"created_at"`
}

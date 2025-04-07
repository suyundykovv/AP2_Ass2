package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

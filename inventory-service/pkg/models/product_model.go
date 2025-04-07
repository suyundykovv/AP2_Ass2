package models

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"category_id"` // вместо Category string
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

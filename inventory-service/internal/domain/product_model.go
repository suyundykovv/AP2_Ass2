package domain

type Product struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"` // вместо Category string
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// type Discount struct {
// 	ID                 string    `json:"id"`
// 	Name               string    `json:"name"`
// 	Description        string    `json:"description"`
// 	DiscountPercentage float64   `json:"discount_percentage"`
// 	ApplicableProducts []int     `json:"applicable_products"` // assuming product IDs
// 	StartDate          time.Time `json:"start_date"`
// 	EndDate            time.Time `json:"end_date"`
// 	IsActive           bool      `json:"is_active"`
// }

package models

type Order struct {
	ID     int32    `json:"id"`
	UserID int      `json:"user_id"`
	Items  []string `json:"items"`
	Status string   `json:"status"`
	Total  float64  `json:"total"`
}

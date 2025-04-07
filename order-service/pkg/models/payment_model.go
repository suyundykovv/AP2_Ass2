package models

type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
	Status  string  `json:"status"`
}

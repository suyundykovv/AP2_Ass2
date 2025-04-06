package models

import "time"

type Order struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Items   []string  `json:"items"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
}

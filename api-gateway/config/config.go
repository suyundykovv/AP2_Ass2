package config

import (
	"os"
)

type Config struct {
	OrderServiceURL     string
	InventoryServiceURL string
}

func LoadConfig() Config {
	return Config{
		OrderServiceURL:     os.Getenv("http://localhost:8082"), // e.g., http://order-service:8082
		InventoryServiceURL: os.Getenv("http://localhost:8081"), // e.g., http://inventory-service:8081
	}
}

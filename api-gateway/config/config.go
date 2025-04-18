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
		// Используй имена переменных окружения, например "ORDER_SERVICE_URL"
		OrderServiceURL:     os.Getenv("ORDER_SERVICE_URL"),     // Чтение из переменной окружения
		InventoryServiceURL: os.Getenv("INVENTORY_SERVICE_URL"), // Чтение из переменной окружения
	}
}

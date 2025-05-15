package config

import (
	"os"
)

type Config struct {
	OrderServiceURL      string
	InventoryServiceURL  string
	UserServiceURL       string
	StatisticsServiceURL string
}

func LoadConfig() Config {
	return Config{
		OrderServiceURL:      os.Getenv("ORDER_SERVICE_URL"),
		InventoryServiceURL:  os.Getenv("INVENTORY_SERVICE_URL"),
		UserServiceURL:       os.Getenv("USER_SERVICE_URL"),
		StatisticsServiceURL: os.Getenv("STATISTICS_SERVICE_URL"),
	}
}

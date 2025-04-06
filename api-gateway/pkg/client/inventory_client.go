package client

import (
	"net/http"
)

func FetchInventory() (*http.Response, error) {
	resp, err := http.Get("http://inventory-service:8081/products")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

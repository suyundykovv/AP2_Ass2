package client

import (
	"net/http"
)

func FetchOrders() (*http.Response, error) {
	resp, err := http.Get("http://order-service:8082/orders")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

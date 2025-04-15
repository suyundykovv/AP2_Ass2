package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type InventoryClient struct {
	BaseURL string
}

func NewInventoryClient(baseURL string) *InventoryClient {
	return &InventoryClient{BaseURL: baseURL}
}

func (c *InventoryClient) ForwardRequest(method, path string, body []byte) ([]byte, error) {
	url := "http://inventory-service:8081/" + path

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

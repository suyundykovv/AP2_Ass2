package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type OrdersClient struct {
	BaseURL string
}

func NewOrdersClient(baseURL string) *OrdersClient {
	return &OrdersClient{BaseURL: baseURL}
}

func (c *OrdersClient) ForwardRequest(method, path string, body []byte) ([]byte, error) {
	url := c.BaseURL + path
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

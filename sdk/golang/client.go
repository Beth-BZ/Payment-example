package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) CreatePayment(data PaymentCreate) (*Payment, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize payment: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/v1/payments/", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var payment Payment
	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &payment, nil
}

func (c *Client) GetPayment(id string) (*Payment, error) {
	resp, err := c.HTTPClient.Get(c.BaseURL + "/v1/payments/" + id)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("payment not found: %s", id)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var payment Payment
	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &payment, nil
}

func (c *Client) UpdatePayment(id string, updates PaymentUpdate) (*Payment, error) {
	body, err := json.Marshal(updates)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize updates: %w", err)
	}

	req, err := http.NewRequest("PATCH", c.BaseURL+"/v1/payments/"+id, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var payment Payment
	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &payment, nil
}

func (c *Client) DeletePayment(id string) error {
	req, err := http.NewRequest("DELETE", c.BaseURL+"/v1/payments/"+id, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return fmt.Errorf("payment not found: %s", id)
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("API error: %s", resp.Status)
	}
	return nil
}

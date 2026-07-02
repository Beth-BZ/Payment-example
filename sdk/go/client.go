// sdk/go/client.go

package payment

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
)

type Client struct {
    BaseURL    string
    HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
    return &Client{
        BaseURL:    baseURL,
        HTTPClient: &http.Client{},
    }
}

// Create payment
func (c *Client) CreatePayment(payload map[string]interface{}) (*http.Response, error) {
    body, _ := json.Marshal(payload)
    return c.HTTPClient.Post(c.BaseURL+"/v1/payments", "application/json", bytes.NewBuffer(body))
}

// Get payment
func (c *Client) GetPayment(id string) (*http.Response, error) {
    return c.HTTPClient.Get(fmt.Sprintf("%s/v1/payments/%s", c.BaseURL, id))
}

// Update payment
func (c *Client) UpdatePayment(id string, payload map[string]interface{}) (*http.Response, error) {
    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/payments/%s", c.BaseURL, id), bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    return c.HTTPClient.Do(req)
}

// Delete payment
func (c *Client) DeletePayment(id string) (*http.Response, error) {
    req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/payments/%s", c.BaseURL, id), nil)
    return c.HTTPClient.Do(req)
}

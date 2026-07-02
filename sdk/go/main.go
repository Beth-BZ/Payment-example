package main

import (
	"fmt"
	"io"
	"bytes"
	"encoding/json"
	"net/http"
)

const BaseURL = "http://localhost:8000"

func main() {
	client := &http.Client{}

	// CREATE
	payload := map[string]interface{}{
		"amount":   76.9,
		"currency": "ETB",
		"status":   "pending",
		"method":   "card",
	}
	body, _ := json.Marshal(payload)
	resp, _ := client.Post(BaseURL+"/v1/payments", "application/json", bytes.NewBuffer(body))
	result, _ := io.ReadAll(resp.Body)
	fmt.Println("CREATE:", string(result))
}

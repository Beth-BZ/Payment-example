package main

type Payment struct {
	ID       string  `json:"id"`
	Date     string  `json:"date"`
	Status   string  `json:"status"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type PaymentCreate struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Status   string  `json:"status"`
}

type PaymentUpdate struct {
	Status   *string  `json:"status,omitempty"`
	Amount   *float64 `json:"amount,omitempty"`
	Currency *string  `json:"currency,omitempty"`
}

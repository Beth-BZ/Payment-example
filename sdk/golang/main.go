package main

import (
	"fmt"
	"log"
)

func main() {
	client := NewClient("http://localhost:8000")

	fmt.Println("--- CREATE ---")
	payment, err := client.CreatePayment(PaymentCreate{
		Amount:   76.9,
		Currency: "ETB",
		Status:   "pending",
	})
	if err != nil {
		log.Fatal("Create failed:", err)
	}
	fmt.Printf("Created: %+v\n", payment)

	fmt.Println("\n--- GET ONE ---")
	fetched, err := client.GetPayment(payment.ID)
	if err != nil {
		log.Fatal("Get failed:", err)
	}
	fmt.Printf("Fetched: %+v\n", fetched)

	fmt.Println("\n--- UPDATE ---")
	status := "completed"
	updated, err := client.UpdatePayment(payment.ID, PaymentUpdate{
		Status: &status,
	})
	if err != nil {
		log.Fatal("Update failed:", err)
	}
	fmt.Printf("Updated: %+v\n", updated)

	fmt.Println("\n--- DELETE ---")
	err = client.DeletePayment(payment.ID)
	if err != nil {
		log.Fatal("Delete failed:", err)
	}
	fmt.Printf("Payment %s deleted successfully\n", payment.ID)
}

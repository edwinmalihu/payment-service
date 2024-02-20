package response

import "time"

type ResponseSuccess struct {
	PaymentID   uint      `json:"payment_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}

type ResponseDelete struct {
	Status string `json:"status"`
}

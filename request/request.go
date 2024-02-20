package request

import "time"

type RequestAddPayment struct {
	OrderID     uint      `json:"order_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}

type RequestByIdPayment struct {
	PaymentID string `form:"json:payment_id"`
}

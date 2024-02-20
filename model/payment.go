package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
}

func (Customer) TableName() string {
	return "customer"
}

type Order struct {
	gorm.Model
	CustomerID  uint      `json:"customer_id"`
	Customer    Customer  `gorm:"foreignKey:CustomerID"`
	TotalAmount float64   `json:"total_amount" gorm:"type:decimal(22,2)"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
}

func (Order) TableName() string {
	return "order"
}

type Payment struct {
	gorm.Model
	OrderID     uint      `json:"order_id"`
	Order       Order     `gorm:"foreignKey:OrderID"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Status      string    `json:"status"`
}

func (Payment) TableName() string {
	return "payments"
}

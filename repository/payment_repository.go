package repository

import (
	"log"
	"payment-service/model"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	Migrate() error
}

type paymentRepo struct {
	DB *gorm.DB
}

// Migrate implements PaymentRepo.
func (p paymentRepo) Migrate() error {
	log.Println("Tabel Payment Migrate Start")
	return p.DB.AutoMigrate(&model.Payment{})
}

func NewPaymentRepo(db *gorm.DB) PaymentRepo {
	return paymentRepo{
		DB: db,
	}
}

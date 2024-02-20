package repository

import (
	"log"
	"payment-service/model"
	"payment-service/request"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	Migrate() error
	AddPayment(request.RequestAddPayment) (model.Payment, error)
	UpdateStatusOrder(uint, string) (model.Order, error)
	DeletePayment(uint) (model.Payment, error)
	DetailPayment(uint) (model.Payment, error)
}

type paymentRepo struct {
	DB *gorm.DB
}

// DetailPayment implements PaymentRepo.
func (p paymentRepo) DetailPayment(id uint) (data model.Payment, errr error) {
	return data, p.DB.First(&data, "id = ?", id).Error
}

// DeletePayment implements PaymentRepo.
func (p paymentRepo) DeletePayment(id uint) (data model.Payment, err error) {
	p.DB.Model(&data).Where("id = ?", id).Update("status", "payment cancel")
	return data, p.DB.Where("id = ?", id).Delete(&data).Error
}

// UpdateStatusOrder implements PaymentRepo.
func (p paymentRepo) UpdateStatusOrder(id_order uint, status string) (data model.Order, err error) {
	return data, p.DB.Model(&data).Where("id = ?", id_order).Update("status", status).Error
}

// AddPayment implements PaymentRepo.
func (p paymentRepo) AddPayment(req request.RequestAddPayment) (data model.Payment, err error) {
	data = model.Payment{
		OrderID:     req.OrderID,
		Amount:      req.Amount,
		PaymentDate: req.PaymentDate,
		Status:      "payment success",
	}

	return data, p.DB.Create(&data).Error
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

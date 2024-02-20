package controller

import (
	"log"
	"net/http"
	"payment-service/repository"
	"payment-service/request"
	"payment-service/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	AddPayment(*gin.Context)
	DeletePayment(*gin.Context)
}

type paymentController struct {
	paymentRepo repository.PaymentRepo
}

// DeletePayment implements PaymentController.
func (p paymentController) DeletePayment(ctx *gin.Context) {
	var req request.RequestByIdPayment
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(req.PaymentID)

	data, err := p.paymentRepo.DetailPayment(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	updateOrder, _ := p.paymentRepo.UpdateStatusOrder(data.OrderID, "payment cancel")
	log.Println(updateOrder)

	delete, err := p.paymentRepo.DeletePayment(data.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	log.Println(delete)

	res := response.ResponseDelete{
		Status: "id payment" + req.PaymentID + "success deleted",
	}

	ctx.JSON(http.StatusOK, res)
}

// AddPayment implements PaymentController.
func (p paymentController) AddPayment(ctx *gin.Context) {
	var req request.RequestAddPayment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := p.paymentRepo.AddPayment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderupdate, err := p.paymentRepo.UpdateStatusOrder(req.OrderID, "payment success")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("orderupdate : ", orderupdate)

	res := response.ResponseSuccess{
		PaymentID:   data.ID,
		Amount:      data.Amount,
		PaymentDate: data.PaymentDate,
		Status:      "Pembayaran Success",
	}

	ctx.JSON(http.StatusOK, res)
}

func NewPaymentController(repo repository.PaymentRepo) PaymentController {
	return paymentController{
		paymentRepo: repo,
	}
}

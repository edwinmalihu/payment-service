package route

import (
	"fmt"
	"log"
	"os"
	"payment-service/controller"
	"payment-service/middleware"
	"payment-service/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB) {
	httpRoute := gin.Default()
	httpRoute.Use(middleware.CORSMiddleware())

	paymentRepo := repository.NewPaymentRepo(db)

	if err := paymentRepo.Migrate(); err != nil {
		log.Fatal("paymet migrate error : ", err)
	}

	paymentController := controller.NewPaymentController(paymentRepo)

	apiRoute := httpRoute.Group("/api")
	{
		apiRoute.POST("/add", paymentController.AddPayment)
		apiRoute.DELETE("/delete", paymentController.DeletePayment)
	}

	//httpRoute.Run(":8088")
	httpRoute.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}

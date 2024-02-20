package route

import (
	"log"
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

	httpRoute.Run(":8084")
}

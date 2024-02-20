package main

import (
	"payment-service/model"
	"payment-service/route"
)

func main() {
	DB, _ := model.DBConnection()
	route.SetupRoute(DB)
}

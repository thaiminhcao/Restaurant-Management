package routes

import (
	controller "Golang-RestaurantManagement/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/Invoice_id", UpdateInvoice())
}
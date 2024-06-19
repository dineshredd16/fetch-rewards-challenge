package routes

import (
	"fetch-rewards-receipt-processor-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	receiptRoutes := router.Group("/receipts")
	{
		receiptRoutes.POST("/process", controllers.ProcessReceipts)
		receiptRoutes.GET("", controllers.GetReceipts)
		receiptRoutes.GET(":id", controllers.GetReceiptById)
		receiptRoutes.GET(":id/points", controllers.GetPoints)
	}
}
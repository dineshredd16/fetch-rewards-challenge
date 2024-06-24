package controllers

import (
	"fetch-rewards-receipt-processor-challenge/models"
	"fetch-rewards-receipt-processor-challenge/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ProcessReceipts(c *gin.Context) {
	var receipt models.Receipt

	err := c.ShouldBindJSON(&receipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Failure", "Message": "Bad Params", "Data" : "", "Error": err.Error(),
		})
	} else if services.CheckMissingReceiptFields(receipt) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Failure", "Message": "Missing Fields", "Data" : "", "Error": "",
		})
	} else {
		services.AddReceipt(&receipt)
			c.JSON(http.StatusOK, gin.H{
				"Status": "Success", "Message": " Receipt Added Succesfully ", "ID" : receipt.Id, "Error": "",
		})
	}
}

func GetPoints (c *gin.Context) {
	id := c.Param("id")
	models.InMemoryDatabase.Lock()
	defer models.InMemoryDatabase.Unlock()
	receipt, exists := models.InMemoryDatabase.Receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"Status": "Failure", "Message": "Receipt Not Found", "Data" : "", "Error": "",
		})
	} else {
		points := services.CalculatePoints(receipt)
		c.JSON(http.StatusOK, gin.H{
			"Status": "Success", "Message": " Receipt Points Calculated Succesfully ", "Points" : points, "Error": "",
		})
	}
}


func GetReceipts (c *gin.Context) {
	models.InMemoryDatabase.Lock()
	defer models.InMemoryDatabase.Unlock()
	receipts := make([]models.Receipt, 0 , len(models.InMemoryDatabase.Receipts))
	for _, receipt := range models.InMemoryDatabase.Receipts {
		receipts = append(receipts, receipt)
	}
	c.JSON(http.StatusOK, gin.H{
		"Status": "Success", "Message": " Fetched Receipts Succesfully ", "Data" : receipts, "Error": "",
	})
}

func UpdateReceipt (c *gin.Context) {
	var receipt models.Receipt
	err := c.ShouldBindJSON(&receipt)
	if (err != nil || services.CheckMissingReceiptFields(receipt)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Failure", "Message": "Bad Params", "Data" : "", "Error": err.Error(),
		})
	} else if _, exists := models.InMemoryDatabase.Receipts[receipt.Id]; !exists{
		c.JSON(http.StatusNotFound, gin.H{
			"Status": "Failure", "Message": "Receipt ID not present", "Data" : "", "Error": "",
		})
	} else {
		models.InMemoryDatabase.Receipts[receipt.Id] = receipt
		c.JSON(http.StatusOK, gin.H{
			"Status": "Success", "Message": " Updated Receipt Succesfully ", "Data" : receipt, "Error": "",
		})
	}
}

func AttachUserToReceipt () {

}

func GetReceiptById (c *gin.Context) {
	id := c.Param("id")
	models.InMemoryDatabase.Lock()
	defer models.InMemoryDatabase.Unlock()
	receipt, exists := models.InMemoryDatabase.Receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"Status": "Failure", "Message": "Receipt Not Found", "Data" : "", "Error": "",
		})
	} else {
			c.JSON(http.StatusOK, gin.H{
				"Status": "Success", "Message": " Fetched Receipts Succesfully ", "Data" : receipt, "Error": "",
		})
	}
}
package services

import (
	"fetch-rewards-receipt-processor-challenge/models"
	"github.com/google/uuid"
)

func GenerateReceiptID () string {
	id := uuid.New()
	return id.String()
}

func CheckMissingReceiptFields (receipt models.Receipt) bool {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" || len(receipt.Items) == 0 {
		return true
	}
	return false
}

func AddReceipt (receipt *models.Receipt) {
	id := GenerateReceiptID()
	receipt.Id = id
	models.InMemoryDatabase.Lock()
	defer models.InMemoryDatabase.Unlock()
	models.InMemoryDatabase.Receipts[receipt.Id] = *receipt
}
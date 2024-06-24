package services

import (
	"fetch-rewards-receipt-processor-challenge/models"
	"math"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints (receipt models.Receipt) int64{
	points := 0
	
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == float64(int(total)) {
		points += 50
	}

	if err == nil && int(total*100)%25 == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && date.Day()%2 != 0 {
		points += 6
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() > 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return int64(points)
}
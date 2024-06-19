package models

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type Receipt struct {
	Id string `json:"id"`
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"PurchaseTime"`
	Total string `json:"total"`
	Items []Item `json:"items"`
}
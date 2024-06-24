package models

type User struct {
	Id string `json: "id"`
	Name string `json: "name"`
	Age int `json: "age"`
	ReceiptID string `json: "receipt_id"`
}
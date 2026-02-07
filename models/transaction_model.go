package models

import "time"

type Transaction struct {
	ID          uint      `json:"id"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransactionDetail struct {
	ID            uint    `json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Quantity      int     `json:"quantity"`
	SubTotal      float64 `json:"sub_total"` // Price * Quantity
}

type TransactionItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type TransactionRequest struct {
	Items []TransactionItem `json:"items"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (TransactionDetail) TableName() string {
	return "transaction_details"
}

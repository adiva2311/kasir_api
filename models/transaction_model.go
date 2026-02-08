package models

import "time"

type Transaction struct {
	ID          uint                `json:"id"`
	TotalAmount float64             `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details" gorm:"foreignKey:TransactionID;references:ID"`
}

type TransactionDetail struct {
	ID            uint `json:"id"`
	TransactionID uint `json:"transaction_id"`
	ProductID     uint `json:"product_id"`
	// ProductName   string  `json:"product_name"`
	Quantity int     `json:"quantity"`
	SubTotal float64 `json:"sub_total"` // Price * Quantity
}

func (Transaction) TableName() string {
	return "transactions"
}

func (TransactionDetail) TableName() string {
	return "transaction_details"
}

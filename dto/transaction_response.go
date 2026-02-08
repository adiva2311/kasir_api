package dto

import "kasir_api/models"

type TransactionResponse struct {
	ID          uint    `json:"id"`
	TotalAmount float64 `json:"total_amount"`
	CreatedAt   string  `json:"created_at"`
	Details     []any   `json:"details"`
}

func ToTransactionResponse(transaction *models.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:          transaction.ID,
		TotalAmount: transaction.TotalAmount,
		CreatedAt:   transaction.CreatedAt.Format("2006-01-02 15:04:05"),
		Details:     []any{},
	}
}

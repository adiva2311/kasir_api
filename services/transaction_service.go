package services

import (
	"kasir_api/models"
	"kasir_api/repositories"
)

type TransactionService interface {
	CreateTransaction(transactionReq *models.TransactionItem) (*models.Transaction, error)
}

type TransactionServiceImpl struct {
	TransactionRepo repositories.TransactionRepo
}

// CreateTransaction implements TransactionService.
func (t *TransactionServiceImpl) CreateTransaction(transactionReq *models.TransactionItem) (*models.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionService(transactionRepo repositories.TransactionRepo) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepo: transactionRepo,
	}
}

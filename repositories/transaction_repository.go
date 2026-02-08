package repositories

import (
	"kasir_api/models"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transaction *models.Transaction) error
}

type TransactionRepoImpl struct {
	DB *gorm.DB
}

// CreateTransaction implements TransactionRepo.
func (t *TransactionRepoImpl) CreateTransaction(transaction *models.Transaction) error {
	return t.DB.Create(transaction).Error
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &TransactionRepoImpl{
		DB: db,
	}
}

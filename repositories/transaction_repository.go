package repositories

import (
	"kasir_api/models"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transaction []models.Transaction) error
}

type TransactionRepoImpl struct {
	DB *gorm.DB
}

// CreateTransaction implements TransactionRepo.
func (t *TransactionRepoImpl) CreateTransaction(transaction []models.Transaction) error {
	result := t.DB.Create(&transaction).Error
	if result != nil {
		return result
	}

	return nil
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &TransactionRepoImpl{
		DB: db,
	}
}

package controllers

import (
	"kasir_api/repositories"
	"kasir_api/services"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type TransactionController interface {
	CreateTransaction(c *echo.Context) error
}

type TransactionControllerImpl struct {
	TransactionService services.TransactionService
}

// CreateTransaction implements TransactionController.
func (t *TransactionControllerImpl) CreateTransaction(c *echo.Context) error {
	panic("unimplemented")
}

func NewTransactionController(db *gorm.DB) TransactionController {
	transactionService := services.NewTransactionService(repositories.NewTransactionRepo(db))
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}

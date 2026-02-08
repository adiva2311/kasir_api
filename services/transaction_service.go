package services

import (
	"fmt"
	"kasir_api/dto"
	"kasir_api/models"
	"kasir_api/repositories"
)

type TransactionService interface {
	CreateTransaction(transactionReq dto.CreateTransactionReq) (*models.Transaction, error)
}

type TransactionServiceImpl struct {
	TransactionRepo repositories.TransactionRepo
}

// CreateTransaction implements TransactionService.
func (t *TransactionServiceImpl) CreateTransaction(transactionReq dto.CreateTransactionReq) (*models.Transaction, error) {
	if len(transactionReq.Items) == 0 {
		return nil, fmt.Errorf("empty transaction request")
	}

	// take first request payload (client sends items in the first element)
	req := transactionReq

	// try to access DB from the concrete repo impl to query products and update stock
	repoImpl, ok := t.TransactionRepo.(*repositories.TransactionRepoImpl)
	if !ok {
		return nil, fmt.Errorf("transaction repo does not expose DB")
	}
	db := repoImpl.DB

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var total float64
	details := make([]models.TransactionDetail, 0, len(req.Items))

	for _, it := range req.Items {
		var product models.Product
		if err := tx.Where("id = ?", it.ProductID).First(&product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if product.Stock < it.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("not enough stock for product %d", it.ProductID)
		}

		sub := float64(product.Price) * float64(it.Quantity)
		total += sub

		details = append(details, models.TransactionDetail{
			ProductID:   uint(product.ID),
			// ProductName: product.Name,
			Quantity:    it.Quantity,
			SubTotal:    sub,
		})

		// decrement product stock
		newStock := product.Stock - it.Quantity
		if err := tx.Model(&product).Update("stock", newStock).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// create transaction (parent), then details (children) so foreign keys are set
	transaction := models.Transaction{TotalAmount: total}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// attach transaction id to details and persist them
	for i := range details {
		details[i].TransactionID = transaction.ID
	}

	if len(details) > 0 {
		if err := tx.Create(&details).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// reload transaction with details
	if err := tx.Preload("Details").First(&transaction, transaction.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &transaction, nil
}

func NewTransactionService(transactionRepo repositories.TransactionRepo) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepo: transactionRepo,
	}
}

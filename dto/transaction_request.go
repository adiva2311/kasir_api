package dto

type TransactionReq struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateTransactionReq struct {
	Items []TransactionReq `json:"items"`
}

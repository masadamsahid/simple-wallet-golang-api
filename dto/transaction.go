package dto

type TransactionRequest struct {
	Amount int64 `json:"amount" validate:"required,gt=0"`
}

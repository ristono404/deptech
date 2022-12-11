package event

type TransactionCreated struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	Qty       uint   `json:"qty" validate:"required"`
}

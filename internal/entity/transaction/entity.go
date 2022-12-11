package category

import (
	"time"

	"github.com/deptech/internal/entity/product"
)

type Entity struct {
	ID           uint64          `json:"id"`
	Type         uint            `json:"type"`
	CreatedAt    time.Time       `json:"created_at"`
	EntityDetail []*EntityDetail `gorm:"foreignKey:TransactionID;references:ID" json:"transaction_details"`
}

func (e *Entity) TableName() string {
	return "transactions"
}

type EntityDetail struct {
	ID            uint64 `json:"id"`
	TransactionID uint64 `json:"transaction_id"`
	ProductID     uint64 `json:"product_id"`
	Qty           uint   `json:"qty"`
	Product       product.Entity
}

func (e *EntityDetail) TableName() string {
	return "transaction_details"
}

type TransactionHistory struct {
	ID        uint64              `json:"id"`
	Type      string              `json:"type"`
	CreatedAt time.Time           `json:"created_at"`
	Detail    []TransactionDetail `json:"details"`
}

type TransactionDetail struct {
	Product string `json:"product"`
	Qty     uint   `json:"qty"`
}

func TransformToList(data []Entity) (res []TransactionHistory) {
	for _, v := range data {
		transaction := TransactionHistory{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
		}
		if v.Type == 1 {
			transaction.Type = "In"
		} else {
			transaction.Type = "Out"
		}

		for _, detail := range v.EntityDetail {
			transDetail := TransactionDetail{
				Product: detail.Product.Name,
				Qty:     detail.Qty,
			}
			transaction.Detail = append(transaction.Detail, transDetail)
		}

		res = append(res, transaction)
	}
	return
}

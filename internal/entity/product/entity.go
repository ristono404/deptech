package product

import (
	"time"

	"github.com/deptech/internal/entity/category"
)

type Entity struct {
	ID          uint64 `json:"id"`
	CategoryID  uint64 `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	Image       string `json:"image"`
	Category    category.Entity
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (e *Entity) TableName() string {
	return "products"
}

type List struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	Image       string `json:"image"`
	Category    string `json:"category"`
}

func TransformToList(data []*Entity) []List {
	result := []List{}
	for _, v := range data {
		res := List{
			Name:        v.Name,
			Description: v.Description,
			Stock:       v.Stock,
			Image:       v.Image,
			Category:    v.Category.Name,
		}
		result = append(result, res)
	}
	return result
}

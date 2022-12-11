package event

type ProductCreated struct {
	ID          uint64 `json:"id"`
	CategoryID  uint64 `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

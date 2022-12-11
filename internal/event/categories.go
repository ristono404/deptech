package event

type CategoryCreated struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

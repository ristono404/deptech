package event

type UserCreated struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required,isodate"`
	Gender    uint   `json:"gender"`
	Password  string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

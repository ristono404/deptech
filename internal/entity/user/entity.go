package user

import (
	"time"
)

type Entity struct {
	ID        uint64     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	BirthDate time.Time  `json:"birth_date"`
	Gender    uint       `json:"gender"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (e *Entity) TableName() string {
	return "users"
}

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expire_at"`
}

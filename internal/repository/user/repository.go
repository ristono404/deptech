package user

import (
	"time"

	userEntity "github.com/ristono404/deptech/internal/entity/user"
	event "github.com/ristono404/deptech/internal/event"
)

type Repository interface {
	Read(id uint64) (*userEntity.Entity, error)
	ReadByEmailPass(event.UserLogin) (*userEntity.Entity, error)
	List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*userEntity.Entity, total uint64, pages uint64, err error)
	SoftDeleteRange(entities *[]userEntity.Entity) error
	FindAllById(ids ...uint64) (*[]userEntity.Entity, error)
	Create(entity *userEntity.Entity) error
	Update(entity *userEntity.Entity) error
}

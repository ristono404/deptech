package user

import (
	"time"

	userEntity "github.com/ristono404/deptech/internal/entity/user"
	event "github.com/ristono404/deptech/internal/event"
)

type Service interface {
	List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*userEntity.Entity, total uint64, pages uint64, err error)
	Create(event.UserCreated) (*userEntity.Entity, error)
	Read(uint64) (*userEntity.Entity, error)
	Login(event.UserLogin) (*userEntity.Entity, error)
	Update(event.UserCreated) (*userEntity.Entity, error)
	SoftDeleteRange(Ids []uint64) error
}

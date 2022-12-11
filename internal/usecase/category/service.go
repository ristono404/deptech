package category

import (
	"time"

	categoryEntity "github.com/ristono404/deptech/internal/entity/category"
	event "github.com/ristono404/deptech/internal/event"
)

type Service interface {
	List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*categoryEntity.Entity, total uint64, pages uint64, err error)
	Create(event.CategoryCreated) (*categoryEntity.Entity, error)
	Read(uint64) (*categoryEntity.Entity, error)
	Update(event.CategoryCreated) (*categoryEntity.Entity, error)
	SoftDeleteRange(Ids []uint64) error
}

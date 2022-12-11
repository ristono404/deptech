package category

import (
	"time"

	categoryEntity "github.com/deptech/internal/entity/category"
)

type Repository interface {
	Read(id uint64) (*categoryEntity.Entity, error)
	List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*categoryEntity.Entity, total uint64, pages uint64, err error)
	SoftDeleteRange(entities *[]categoryEntity.Entity) error
	FindAllById(ids ...uint64) (*[]categoryEntity.Entity, error)
	Create(entity *categoryEntity.Entity) error
	Update(entity *categoryEntity.Entity) error
}

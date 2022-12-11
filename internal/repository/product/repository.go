package product

import (
	"time"

	productEntity "github.com/deptech/internal/entity/product"
	transEntity "github.com/deptech/internal/entity/transaction"
)

type Repository interface {
	Read(id uint64) (*productEntity.Entity, error)
	List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*productEntity.Entity, total uint64, pages uint64, err error)
	SoftDeleteRange(entities *[]productEntity.Entity) error
	FindAllById(ids ...uint64) (*[]productEntity.Entity, error)
	Create(*productEntity.Entity) error
	Update(*productEntity.Entity) error
	CreateTransaction(*transEntity.Entity) error
	ListTransaction() ([]transEntity.Entity, error)
}

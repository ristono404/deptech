package product

import (
	"time"

	productEntity "github.com/deptech/internal/entity/product"
	transEntity "github.com/deptech/internal/entity/transaction"
	event "github.com/deptech/internal/event"
)

type Service interface {
	List(offset, limit int, from, to *time.Time, search string, by []string) (entries []productEntity.List, total uint64, pages uint64, err error)
	Create(event.ProductCreated) (*productEntity.Entity, error)
	Read(uint64) (*productEntity.Entity, error)
	Update(event.ProductCreated) (*productEntity.Entity, error)
	SoftDeleteRange(Ids []uint64) error
	In([]event.TransactionCreated) (*transEntity.Entity, error)
	Out([]event.TransactionCreated) (*transEntity.Entity, error)
	TransactionHistory() ([]transEntity.TransactionHistory, error)
}

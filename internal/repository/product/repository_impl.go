package product

import (
	"errors"
	"time"

	productEntity "github.com/deptech/internal/entity/product"
	transEntity "github.com/deptech/internal/entity/transaction"
	pkgRepository "github.com/deptech/internal/pkg/repository"
	"github.com/deptech/internal/shared/database"
	"gorm.io/gorm"
)

type repository struct {
	db     *database.Database
	helper pkgRepository.Repository
}

func New(db *database.Database) *repository {
	return &repository{
		db:     db,
		helper: *pkgRepository.New(db.DB),
	}
}

func (r *repository) Read(id uint64) (*productEntity.Entity, error) {
	entry := productEntity.Entity{}
	result := r.db.DB.Where("id = ? AND deleted_at IS NULL", id).First(&entry)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &entry, nil
}

func (r *repository) List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*productEntity.Entity, total, pages uint64, err error) {
	entity := productEntity.Entity{}
	entityTb := entity.TableName()

	for i, v := range by {
		if v == "category" {
			by[i] = "ct.name"
		} else {
			by[i] = "products." + v
		}
	}

	tx, total, pages := r.helper.List(
		offset,
		limit,
		from,
		to,
		search,
		by,
		[]string{"products.id", "products.name", "products.description", "ct.name"},
		entityTb,
		"inner join product_categories ct ON ct.id=products.category_id",
	)

	err = tx.Debug().Preload("Category").Where("products.deleted_at is null").Offset(offset).Limit(limit).Find(&entries).Error

	return
}

func (r *repository) SoftDeleteRange(entities *[]productEntity.Entity) error {
	tx := r.db.DB.Begin()
	if len(*entities) > 0 {
		for _, entity := range *entities {
			now := time.Now()
			entity.DeletedAt = &now
			if err := tx.Where("id = ?", entity.ID).Updates(&entity).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

func (r *repository) FindAllById(ids ...uint64) (*[]productEntity.Entity, error) {
	result := &[]productEntity.Entity{}
	if len(ids) <= 0 {
		return result, nil
	}

	queryRes := r.db.DB.Where("id IN (?) AND deleted_at is null", ids).Find(result)

	if err := queryRes.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Create(entity *productEntity.Entity) error {
	if err := r.db.DB.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(entity *productEntity.Entity) error {
	entity.UpdatedAt = time.Now()
	if err := r.db.DB.Save(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateTransaction(entity *transEntity.Entity) error {
	tx := r.db.DB.Begin()
	if err := tx.Omit("EntityDetail").Create(entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	for i, _ := range entity.EntityDetail {
		entity.EntityDetail[i].TransactionID = entity.ID
	}
	if err := r.createDetail(tx, entity.Type, entity.EntityDetail); err != nil {
		tx.Rollback()
		return err
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
	}
	return nil
}

func (r *repository) createDetail(tx *gorm.DB, typ uint, entities []*transEntity.EntityDetail) error {
	for _, entity := range entities {

		if err := tx.Create(entity).Error; err != nil {
			return err
		}
		productEntity, err := r.Read(entity.ProductID)
		if productEntity == nil {
			err = errors.New("product not found")
		}
		if err != nil {
			return err
		}
		productEntity.UpdatedAt = time.Now()
		if typ == 0 {
			productEntity.Stock -= entity.Qty
		} else {
			productEntity.Stock += entity.Qty
		}

		if err := tx.Where("id = ?", entity.ProductID).Updates(&productEntity).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) ListTransaction() (entries []transEntity.Entity, err error) {
	err = r.db.DB.Debug().Preload("EntityDetail.Product").Find(&entries).Error

	return
}

package category

import (
	"time"

	categoryEntity "github.com/deptech/internal/entity/category"
	pkgRepository "github.com/deptech/internal/pkg/repository"
	"github.com/deptech/internal/shared/database"
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

func (r *repository) Read(id uint64) (*categoryEntity.Entity, error) {
	entry := categoryEntity.Entity{}
	result := r.db.DB.Where("id = ? AND deleted_at IS NULL", id).First(&entry)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &entry, nil
}

func (r *repository) List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*categoryEntity.Entity, total, pages uint64, err error) {
	entity := categoryEntity.Entity{}
	entityTb := entity.TableName()
	tx, total, pages := r.helper.List(
		offset,
		limit,
		from,
		to,
		search,
		by,
		[]string{"id", "name", "description"},
		entityTb,
		"",
	)

	err = tx.Debug().Offset(offset).Limit(limit).Find(&entries).Error

	return
}

func (r *repository) SoftDeleteRange(entities *[]categoryEntity.Entity) error {
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

func (r *repository) FindAllById(ids ...uint64) (*[]categoryEntity.Entity, error) {
	result := &[]categoryEntity.Entity{}
	if len(ids) <= 0 {
		return result, nil
	}

	queryRes := r.db.DB.Where("id IN (?) AND deleted_at is null", ids).Find(result)

	if err := queryRes.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Create(entity *categoryEntity.Entity) error {
	if err := r.db.DB.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(entity *categoryEntity.Entity) error {
	entity.UpdatedAt = time.Now()
	if err := r.db.DB.Save(&entity).Error; err != nil {
		return err
	}
	return nil
}

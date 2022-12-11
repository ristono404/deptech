package user

import (
	"time"

	userEntity "github.com/ristono404/deptech/internal/entity/user"
	"github.com/ristono404/deptech/internal/event"
	pkgRepository "github.com/ristono404/deptech/internal/pkg/repository"
	"github.com/ristono404/deptech/internal/shared/database"
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

func (r *repository) Read(id uint64) (*userEntity.Entity, error) {
	entry := userEntity.Entity{}
	result := r.db.DB.Where("id = ? AND deleted_at IS NULL", id).First(&entry)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &entry, nil
}

func (r *repository) ReadByEmailPass(param event.UserLogin) (*userEntity.Entity, error) {
	entry := userEntity.Entity{}
	result := r.db.DB.Where("email = ? AND password = ? AND deleted_at IS NULL", param.Email, param.Password).First(&entry)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &entry, nil
}

func (r *repository) List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*userEntity.Entity, total, pages uint64, err error) {
	entity := userEntity.Entity{}
	entityTb := entity.TableName()
	tx, total, pages := r.helper.List(
		offset,
		limit,
		from,
		to,
		search,
		by,
		[]string{"id", "first_name", "last_name", "email", "gender"},
		entityTb,
		"",
	)

	err = tx.Debug().Offset(offset).Limit(limit).Find(&entries).Error

	return
}

func (r *repository) SoftDeleteRange(entities *[]userEntity.Entity) error {
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

func (r *repository) FindAllById(ids ...uint64) (*[]userEntity.Entity, error) {
	result := &[]userEntity.Entity{}
	if len(ids) <= 0 {
		return result, nil
	}

	queryRes := r.db.DB.Where("id IN (?) AND deleted_at is null", ids).Find(result)

	if err := queryRes.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Create(entity *userEntity.Entity) error {
	if err := r.db.DB.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(entity *userEntity.Entity) error {
	entity.UpdatedAt = time.Now()
	if err := r.db.DB.Save(&entity).Error; err != nil {
		return err
	}
	return nil
}

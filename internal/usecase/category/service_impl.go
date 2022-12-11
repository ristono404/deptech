package category

import (
	"fmt"
	"time"

	categoryEntity "github.com/ristono404/deptech/internal/entity/category"
	event "github.com/ristono404/deptech/internal/event"
	categoryRepository "github.com/ristono404/deptech/internal/repository/category"
)

type service struct {
	repo categoryRepository.Repository
}

func New(repo categoryRepository.Repository) *service {
	return &service{repo}
}

func (s *service) List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*categoryEntity.Entity, total uint64, pages uint64, err error) {
	entries, total, pages, err = s.repo.List(offset, limit, from, to, search, by)

	return
}

func (s *service) Create(param event.CategoryCreated) (*categoryEntity.Entity, error) {

	entity := &categoryEntity.Entity{
		Name:        param.Name,
		Description: param.Description,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) Read(id uint64) (*categoryEntity.Entity, error) {
	result, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) Update(param event.CategoryCreated) (*categoryEntity.Entity, error) {

	entity, err := s.Read(param.ID)
	if err != nil {
		return nil, err
	}
	entity.Name = param.Name
	entity.Description = param.Description

	if err := s.repo.Update(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) SoftDeleteRange(Ids []uint64) error {
	// check data
	deletedData, err := s.repo.FindAllById(Ids...)
	if err != nil {
		return err
	}
	if len(*deletedData) == 0 {
		return fmt.Errorf("data not found")
	}

	return s.repo.SoftDeleteRange(deletedData)
}

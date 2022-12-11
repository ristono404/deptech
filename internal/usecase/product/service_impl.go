package product

import (
	"fmt"
	"time"

	productEntity "github.com/deptech/internal/entity/product"
	transEntity "github.com/deptech/internal/entity/transaction"
	"github.com/deptech/internal/event"
	"github.com/deptech/internal/pkg/file"
	productRepository "github.com/deptech/internal/repository/product"
)

type service struct {
	repo productRepository.Repository
}

func New(repo productRepository.Repository) *service {
	return &service{repo}
}

func (s *service) List(offset, limit int, from, to *time.Time, search string, by []string) (list []productEntity.List, total uint64, pages uint64, err error) {
	entries := []*productEntity.Entity{}
	entries, total, pages, err = s.repo.List(offset, limit, from, to, search, by)
	list = productEntity.TransformToList(entries)
	return
}

func (s *service) Create(param event.ProductCreated) (*productEntity.Entity, error) {

	entity := &productEntity.Entity{
		CategoryID:  param.CategoryID,
		Name:        param.Name,
		Description: param.Description,
		Stock:       param.Stock,
	}

	if param.Image != "" {
		fileName, err := file.UploadImage(param.Image, "product")
		if err != nil {
			return nil, err
		}
		entity.Image = fileName
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	return entity, nil
}
func (s *service) Read(id uint64) (*productEntity.Entity, error) {
	result, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) Update(param event.ProductCreated) (*productEntity.Entity, error) {

	entity, err := s.Read(param.ID)
	if err != nil {
		return nil, err
	}
	entity.Name = param.Name
	entity.Description = param.Description
	entity.CategoryID = param.CategoryID
	entity.Stock = param.Stock

	if param.Image != "" {
		fileName, err := file.UploadImage(param.Image, "product")
		if err != nil {
			return nil, err
		}
		entity.Image = fileName
	}

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

func (s *service) In(params []event.TransactionCreated) (*transEntity.Entity, error) {

	details := []*transEntity.EntityDetail{}

	for _, v := range params {
		detail := &transEntity.EntityDetail{
			ProductID: v.ProductID,
			Qty:       v.Qty,
		}
		details = append(details, detail)
	}

	entity := &transEntity.Entity{
		Type:         1,
		EntityDetail: details,
	}

	if err := s.repo.CreateTransaction(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) Out(params []event.TransactionCreated) (*transEntity.Entity, error) {

	details := []*transEntity.EntityDetail{}

	for _, v := range params {
		detail := &transEntity.EntityDetail{
			ProductID: v.ProductID,
			Qty:       v.Qty,
		}
		details = append(details, detail)
	}

	entity := &transEntity.Entity{
		Type:         0,
		EntityDetail: details,
	}

	if err := s.repo.CreateTransaction(entity); err != nil {
		return nil, err
	}

	return entity, nil
}
func (s *service) TransactionHistory() ([]transEntity.TransactionHistory, error) {
	entries, err := s.repo.ListTransaction()
	if err != nil {
		return nil, err
	}
	if entries == nil {
		return nil, nil
	}

	return transEntity.TransformToList(entries), nil
}

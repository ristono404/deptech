package user

import (
	"fmt"
	"time"

	userEntity "github.com/deptech/internal/entity/user"
	event "github.com/deptech/internal/event"
	userRepository "github.com/deptech/internal/repository/user"
)

type service struct {
	repo userRepository.Repository
}

func New(repo userRepository.Repository) *service {
	return &service{repo}
}

func (s *service) List(offset, limit int, from, to *time.Time, search string, by []string) (entries []*userEntity.Entity, total uint64, pages uint64, err error) {
	entries, total, pages, err = s.repo.List(offset, limit, from, to, search, by)

	return
}

func (s *service) Create(param event.UserCreated) (*userEntity.Entity, error) {

	format := fmt.Sprintf("%s 00:00:00", param.BirthDate)
	dateParam, _ := time.Parse("2006-01-02 15:04:05", format)

	entity := &userEntity.Entity{
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Email:     param.Email,
		BirthDate: dateParam,
		Gender:    param.Gender,
		Password:  param.Password,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *service) Read(id uint64) (*userEntity.Entity, error) {
	result, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s *service) Login(param event.UserLogin) (*userEntity.Entity, error) {
	result, err := s.repo.ReadByEmailPass(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) Update(param event.UserCreated) (*userEntity.Entity, error) {

	entity, err := s.Read(param.ID)
	if err != nil {
		return nil, err
	}
	entity.FirstName = param.FirstName
	entity.LastName = param.LastName
	entity.Email = param.Email
	entity.Password = param.Password
	entity.Gender = param.Gender

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

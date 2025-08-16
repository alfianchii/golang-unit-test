package service

import (
	"errors"

	"github.com/alfianchii/golang-unit-test/entity"
	"github.com/alfianchii/golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (s CategoryService) Get(id string) (*entity.Category, error) {
	category := s.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
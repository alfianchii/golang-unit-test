package repository

import (
	"github.com/alfianchii/golang-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (r *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := r.Mock.Called(id)
	return args.Get(0).(*entity.Category)
}
package repository

import "github.com/alfianchii/golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
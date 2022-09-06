package services

import (
	"errors"

	"github.com/tiyan-attirmidzi/learn-go-unit-test/entities"
	"github.com/tiyan-attirmidzi/learn-go-unit-test/repositories"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func (s CategoryService) Get(id string) (*entities.Category, error) {
	category := s.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}
}

package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/tiyan-attirmidzi/learn-go-unit-test/entities"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (r *CategoryRepositoryMock) FindById(id string) *entities.Category {
	arguments := r.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entities.Category)
		return &category
	}
}

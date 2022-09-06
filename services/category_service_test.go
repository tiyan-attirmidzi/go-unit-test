package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tiyan-attirmidzi/learn-go-unit-test/entities"
	"github.com/tiyan-attirmidzi/learn-go-unit-test/repositories"
)

var (
	categoryRepository = &repositories.CategoryRepositoryMock{Mock: mock.Mock{}}
	categoryService    = CategoryService{Repository: categoryRepository}
)

func TestCategoryService_GetNotFound(t *testing.T) {
	// Program Mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entities.Category{
		ID:   "2",
		Name: "PC Geming Spek Romusa",
	}
	// Program Mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}

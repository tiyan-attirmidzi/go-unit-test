package repositories

import "github.com/tiyan-attirmidzi/learn-go-unit-test/entities"

type CategoryRepository interface {
	FindById(id string) *entities.Category
}

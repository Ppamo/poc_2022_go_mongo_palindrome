package repositories

import (
	"entities"
)

type repo struct{}

func NewMongoRepository() ProductsRepository {
	return &repo{}
}

func (*repo) FindAll() ([]entities.Product, error) {
	return []entities.Product{}, nil
}

func (*repo) Find(id string) (entities.Product, error) {
	return entities.Product{}, nil
}

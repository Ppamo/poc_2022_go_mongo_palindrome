package services

import (
	"entities"
	"repositories"
)

var (
	repo repositories.ProductsRepository
)

type ProductsService interface {
	FindAll() ([]entities.Product, error)
	Find(id int) (entities.Product, error)
}

type service struct{}

func NewProductsService(r repositories.ProductsRepository) ProductsService {
	repo = r
	return &service{}
}

func (*service) FindAll() ([]entities.Product, error) {
	return repo.FindAll()
}

func (*service) Find(id int) (entities.Product, error) {
	return repo.Find(id)
}

package services

import (
	"entities"
	"repositories"
)

var (
	repo repositories.ProductsRepository
)

type ProductsService interface {
	Find(id int) (entities.Product, error)
	FindText(s string) ([]entities.Product, error)
	FindAll() ([]entities.Product, error)
}

type service struct{}

func NewProductsService(r repositories.ProductsRepository) ProductsService {
	repo = r
	return &service{}
}

func (*service) FindText(s string) ([]entities.Product, error) {
	return repo.FindText(s)
}

func (*service) Find(id int) (entities.Product, error) {
	return repo.Find(id)
}

func (*service) FindAll() ([]entities.Product, error) {
	return repo.FindAll()
}

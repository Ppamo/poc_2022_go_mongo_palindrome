package repositories

import "entities"

type ProductsRepository interface {
	FindAll() ([]entities.Product, error)
	Find(id string) (entities.Product, error)
}

package repositories

import "entities"

type ProductsRepository interface {
	FindAll() ([]entities.Product, error)
	Find(id int) (entities.Product, error)
}

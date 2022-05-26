package repositories

import "entities"

type ProductsRepository interface {
	FindText(s string) ([]entities.Product, error)
	Find(id int) (entities.Product, error)
	FindAll() ([]entities.Product, error)
}

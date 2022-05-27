package repositories

import (
	"entities"
)

type mockrepo struct{}

func NewMockRepository() ProductsRepository {
	return &mockrepo{}
}

func (*mockrepo) FindText(s string) ([]entities.Product, error) {
	products := []entities.Product{
		{ID: 1, Brand: "ooy eqrceli", Description: "rlñlw brhrka", Image: "www.lider.cl/catalogo/images/whiteLineIcon1.svg", Price: 10000},
		{ID: 2, Brand: "azwan ubdehk", Description: "azwan ubdehk", Image: "www.lider.cl/catalogo/images/whiteLineIcon2.svg", Price: 20000},
		{ID: 3, Brand: "frky givgxzqk", Description: "frky givgxzqk", Image: "www.lider.cl/catalogo/images/whiteLineIcon3.svg", Price: 30000},
	}
	return products, nil
}

func (*mockrepo) Find(id int) (entities.Product, error) {
	product := entities.Product{ID: 1, Brand: "ooy eqrceli", Description: "rlñlw brhrka", Image: "www.lider.cl/catalogo/images/whiteLineIcon.svg", Price: 498724}
	return product, nil
}

func (*mockrepo) FindAll() ([]entities.Product, error) {
	products := []entities.Product{
		{ID: 4, Brand: "ooy eqrceli", Description: "rlñlw brhrka", Image: "www.lider.cl/catalogo/images/whiteLineIcon1.svg", Price: 10000},
		{ID: 5, Brand: "azwan ubdehk", Description: "azwan ubdehk", Image: "www.lider.cl/catalogo/images/whiteLineIcon2.svg", Price: 20000},
		{ID: 6, Brand: "frky givgxzqk", Description: "frky givgxzqk", Image: "www.lider.cl/catalogo/images/whiteLineIcon3.svg", Price: 30000},
	}
	return products, nil
}

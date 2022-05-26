package controllers

import (
	ce "customerrors"
	"encoding/json"
	"net/http"
	"services"
)

type controller struct{}

var (
	productsService services.ProductsService
)

type ProductsController interface {
	GetProduct(response http.ResponseWriter, request *http.Request)
	GetProducts(response http.ResponseWriter, request *http.Request)
}

func NewProductsController(service services.ProductsService) ProductsController {
	productsService = service
	return &controller{}
}

func (*controller) GetProducts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	products, err := productsService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ce.ServiceError{Message: "Error getting products"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(products)
}

func (*controller) GetProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	id := "1"
	product, err := productsService.Find(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ce.ServiceError{Message: "Error getting the product"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(product)
}

package controllers

import (
	ce "customerrors"
	"encoding/json"
	"entities"
	"net/http"
	"services"
	"strconv"
	"strings"
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

func checkPalindrome(w string) bool {
	j := len(w) - 1
	for i := 0; i < (len(w) / 2); i++ {
		if w[i] != w[j] {
			return false
		}
		j--
	}
	return true
}

func (*controller) GetProducts(response http.ResponseWriter, request *http.Request) {
	var (
		products []entities.Product
		err      error
	)
	response.Header().Set("Content-Type", "application/json")
	q := request.URL.Query().Get("q")
	if len(q) > 0 {
		products, err = productsService.FindText(q)
		if checkPalindrome(q) {
			for i, _ := range products {
				products[i].Discount = 50
			}
		}
	} else {
		products, err = productsService.FindAll()
	}
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
	id, _ := strconv.Atoi(strings.Replace(request.URL.Path, "/product/", "", 1))
	product, err := productsService.Find(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(ce.ServiceError{Message: "Error getting the product"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(product)
}

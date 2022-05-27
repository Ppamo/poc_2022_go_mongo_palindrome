package controllers

import (
	"encoding/json"
	"entities"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"repositories"
	"services"
	"testing"
)

var (
	m repositories.ProductsRepository = repositories.NewMockRepository()
	s services.ProductsService        = services.NewProductsService(m)
	c ProductsController              = NewProductsController(s)
)

func TestPalindrome(t *testing.T) {
	var (
		non_palindrome = []string{"asd", "asda"}
		palindrome     = []string{"abba", "saas", "saaas", "baa1aab"}
	)
	for _, item := range non_palindrome {
		assert.False(t, checkPalindrome(item), "Palindrome should have been true")
	}
	for _, item := range palindrome {
		assert.True(t, checkPalindrome(item), "Palindrome should have been false")
	}
}

func TestGetProducts(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000/products", nil)
	w := httptest.NewRecorder()
	c.GetProducts(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status Code should have been 200")
	products := []entities.Product{}
	json.Unmarshal(body, &products)
	assert.Equal(t, len(products), 3, "Controller should have been returned 3 products")
	assert.Equal(t, products[0].Discount, 0, "Controller should have been returned products with 0 discount")
}

func TestGetProductsWithText(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000/products?q=saa", nil)
	w := httptest.NewRecorder()
	c.GetProducts(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status Code should have been 200")
	products := []entities.Product{}
	json.Unmarshal(body, &products)
	assert.Equal(t, len(products), 3, "Controller should have been returned 3 products")
	assert.Equal(t, products[0].Discount, 0, "Controller should have been returned products with 0 discount")
}

func TestGetProductsWithTextPalindrome(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000/products?q=saas", nil)
	w := httptest.NewRecorder()
	c.GetProducts(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status Code should have been 200")
	products := []entities.Product{}
	json.Unmarshal(body, &products)
	assert.Equal(t, len(products), 3, "Controller should have been returned 3 products")
	assert.Equal(t, products[0].Discount, 50, "Controller should have been returned products with 50 discount")
}

func TestGetProduct(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:3000/product/1", nil)
	w := httptest.NewRecorder()
	c.GetProduct(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status Code should have been 200")
	product := entities.Product{}
	json.Unmarshal(body, &product)
	assert.Equal(t, product.ID, 1, "Controller should have been returned product 1")
	assert.Equal(t, product.Discount, 0, "Controller should have been returned a product with 0 discount")
}

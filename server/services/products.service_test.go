package services

import (
	"github.com/stretchr/testify/assert"
	"os"
	"repositories"
	"testing"
)

var (
	m repositories.ProductsRepository = repositories.NewMockRepository()
	s ProductsService                 = NewProductsService(m)
)

func TestMain(m *testing.M) {
	os.Setenv("MODE", "TEST")
	os.Exit(m.Run())
}

func TestServiceFindId(t *testing.T) {
	p, err := s.Find(1)
	assert.Equal(t, p.ID, 1, "Brand should have been 1")
	assert.Equal(t, p.Brand, "ooy eqrceli", "Brand should have been 'ooy eqrceli'")
	assert.Equal(t, p.Description, "rlñlw brhrka", "Brand should have been 'rlñlw brhrka'")
	assert.Equal(t, p.Image, "www.lider.cl/catalogo/images/whiteLineIcon.svg", "Brand should have been 'alogo/images/whiteLineIcon.svg'")
	assert.Equal(t, p.Price, int64(498724), "Brand should have been 498724")
	assert.Equal(t, p.Discount, 0, "Brand should have been 0")
	assert.Nil(t, err, "Error should have been nil")
}

func TestServiceFindText(t *testing.T) {
	p, err := s.FindText("")
	assert.Nil(t, err, "Error should have been nil")
	assert.Equal(t, len(p), 3, "Len should have been 3")
	assert.Equal(t, p[0].ID, 1, "Product[0's ID should have been 1")
	assert.Equal(t, p[1].ID, 2, "Product[0's ID should have been 2")
	assert.Equal(t, p[2].ID, 3, "Product[0's ID should have been 3")
}

func TestServiceFindAll(t *testing.T) {
	p, err := s.FindAll()
	assert.Nil(t, err, "Error should have been nil")
	assert.Equal(t, len(p), 3, "Len should have been 3")
	assert.Equal(t, p[0].ID, 4, "Product[0's ID should have been 4")
	assert.Equal(t, p[1].ID, 5, "Product[0's ID should have been 5")
	assert.Equal(t, p[2].ID, 6, "Product[0's ID should have been 6")
}

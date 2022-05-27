package main

import (
	"controllers"
	"log"
	"os"
	"repositories"
	"routers"
	"services"
)

const staticPath = "static"

var (
	r router.Router
	m repositories.ProductsRepository
	s services.ProductsService
	c controllers.ProductsController
)

func setup() {
	r = router.NewMuxRouter()
	if os.Getenv("MODE") == "TEST" {
		m = repositories.NewMockRepository()
	} else {
		m = repositories.NewMongoRepository()
	}
	s = services.NewProductsService(m)
	c = controllers.NewProductsController(s)
}

func main() {
	setup()
	log.Printf("> Starting server")
	port := os.Getenv("PORT")
	r.GET("/products", c.GetProducts)
	r.GET("/product/{id}", c.GetProduct)
	r.STATIC("/", "./static/")
	r.SERVE(port)
}

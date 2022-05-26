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
	r router.Router                   = router.NewMuxRouter()
	m repositories.ProductsRepository = repositories.NewMongoRepository()
	s services.ProductsService        = services.NewProductsService(m)
	c controllers.ProductsController  = controllers.NewProductsController(s)
)

func main() {
	log.Printf("> Starting server")
	port := os.Getenv("PORT")
	r.GET("/products", c.GetProducts)
	r.GET("/product/{id}", c.GetProduct)
	r.STATIC("/", "./static/")
	r.SERVE(port)
}

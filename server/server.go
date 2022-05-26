package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"routers"
)

const staticPath = "static"

var (
	r router.Router = router.NewMuxRouter()
)

func handlerAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API")
}

func main() {
	log.Printf("> Starting server")
	port := os.Getenv("PORT")
	r.GET("/", handlerAPI)
	r.GET("/product/:id", handlerAPI)
	r.GET("/products", handlerAPI)
	r.STATIC("/", "./static/")
	r.SERVE(port)
}

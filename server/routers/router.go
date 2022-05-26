package router

import (
	"net/http"
)

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	STATIC(uri string, path string)
	SERVE(port string)
}

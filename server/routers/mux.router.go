package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxRouter struct{}

var muxDispatcher = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) SERVE(port string) {
	var address string = fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("> Mux router listening at %s", address)
	err := http.ListenAndServe(address, muxDispatcher)
	if err != nil {
		log.Println(err)
	}
}

func (*muxRouter) STATIC(uri string, path string) {
	muxDispatcher.PathPrefix(uri).Handler(http.FileServer(http.Dir(path)))
}

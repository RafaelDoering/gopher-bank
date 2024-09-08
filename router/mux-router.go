package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

type muxRouter struct { }

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) Serve(port string) {
	muxPort := ":" + port

	fmt.Printf("Mux server listening on port: %v", muxPort)
	http.ListenAndServe(muxPort, router)
}

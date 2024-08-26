package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Serve(port string) {
	muxPort := ":" + port

	router := mux.NewRouter()
	fmt.Printf("Mux server listening on port: %v", muxPort)
	http.ListenAndServe(muxPort, router)
}

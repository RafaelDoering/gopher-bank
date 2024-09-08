package router

import "net/http"

type Router interface {
	Serve(port string)
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
}

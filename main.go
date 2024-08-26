package main

import (
	"gopher-bank/router"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = "8080"

	httpRouter.Serve(port)
}